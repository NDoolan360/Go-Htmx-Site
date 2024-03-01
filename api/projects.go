package api

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	githublangsgo "github.com/NDoolan360/github-langs-go"
)

// GetProjects handles the request for fetching and rendering project data.
func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, errs := FetchProjects(r.URL.Query()["host"])
	if len(projects) > 0 {
		projectTemplate := template.Must(template.ParseFiles(
			GetApiAsset("template/projects.gohtml"),
		))
		if err := projectTemplate.Execute(w, ProjectsTemplate{Projects: projects}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
	}
}

func FetchProjects(hostNames []string) (projects []Project, errs error) {
	for _, hostName := range hostNames {
		host, ok := hostMap[hostName]
		if ok {
			data, err := Fetch(&host.Request)
			if err != nil {
				errs = errors.Join(errs, err)
			}
			newProjects, err := host.Parser(data)
			if err != nil {
				errs = errors.Join(errs, err)
			}
			projects = append(projects, newProjects...)
		} else {
			errs = errors.Join(errs, fmt.Errorf("URL not found for host: %s", hostName))
		}
	}
	return projects, errs
}

// Fetch fetches content from a given URL using HTTP GET.
var Fetch = func(request *Request) ([]byte, error) {
	body := strings.NewReader(request.Body)
	outgoingRequest, err := http.NewRequest(request.Method, request.Path, body)
	if err != nil {
		return nil, err
	}
	if len(request.ContentType) > 0 {
		outgoingRequest.Header.Add("Content-Type", "application/json")
	}
	if len(request.Username) > 0 && len(request.Password) > 0 {
		outgoingRequest.SetBasicAuth(request.Username, request.Password)
	}

	client := &http.Client{}
	response, err := client.Do(outgoingRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if data, err := io.ReadAll(response.Body); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

var hostMap = map[string]Host{
	"github": {
		Request: Request{
			Method: "GET",
			Path:   "https://api.github.com/users/NDoolan360/repos?sort=stars",
		},
		Parser: func(data []byte) (projects []Project, err error) {
			var githubProjects []struct {
				Title       string   `json:"name"`
				Description string   `json:"description"`
				Url         string   `json:"html_url"`
				Language    string   `json:"language"`
				Topics      []string `json:"topics"`
				Fork        bool     `json:"fork"`
			}

			if unmarshalErr := json.Unmarshal(data, &githubProjects); unmarshalErr != nil {
				return nil, errors.Join(errors.New("error parsing GitHub projects"), unmarshalErr)
			}

			for _, project := range githubProjects {
				if project.Fork || len(project.Topics) == 0 {
					continue
				}

				var colour template.CSS
				if lang, languageErr := githublangsgo.GetLanguage(project.Language); languageErr != nil {
					err = errors.Join(err, fmt.Errorf("error parsing language (%s)", project.Language), languageErr)
				} else {
					colour = template.CSS(lang.Color)
				}

				projects = append(projects, Project{
					Host:  "Github",
					Title: project.Title,
					Url:   project.Url,
					Language: Language{
						Name:   project.Language,
						Colour: colour,
					},
					Logo:        GetSVGLogo("github"),
					Description: template.HTML(project.Description),
					Topics:      project.Topics,
				})
			}

			return projects, err
		},
	},
	"cults3d": {
		Request: Request{
			Method:   "POST",
			Path:     "https://cults3d.com/graphql",
			Username: os.Getenv("CULTS3D_USERNAME"),
			Password: os.Getenv("CULTS3D_API_KEY"),
			Body: `{
				"query": "{
					creationsBatch(limit: 5, creatorNick: "ND360", sort: BY_DOWNLOADS) {
						results {
							name(locale: EN)
							url(locale: EN)
							illustrationImageUrl
							tags
						}
					}
				}"
			}`,
		},
		Parser: func(data []byte) (projects []Project, err error) {
			var cults3dProjects struct {
				Data struct {
					CreationsBatch struct {
						Results []struct {
							Title       string   `json:"name"`
							Description string   `json:"description"`
							Url         string   `json:"url"`
							ImageSrc    string   `json:"illustrationImageUrl"`
							Topics      []string `json:"tags"`
						} `json:"results"`
					} `json:"creationsBatch"`
				} `json:"data"`
			}
			if unmarshalErr := json.Unmarshal(data, &cults3dProjects); unmarshalErr != nil {
				return nil, errors.Join(errors.New("error parsing Cults3D projects"), unmarshalErr)
			}

			for _, project := range cults3dProjects.Data.CreationsBatch.Results {
				projects = append(projects, Project{
					Host:        "Cults3D",
					Title:       project.Title,
					Description: template.HTML(project.Description),
					Url:         project.Url,
					Image: Image{
						Src: template.HTMLAttr(fmt.Sprintf("src=\"%s\"", project.ImageSrc)),
						Alt: template.HTMLAttr(fmt.Sprintf("alt=\"3D Model: %s\"", project.Title)),
					},
					Logo:   GetSVGLogo("cults3d"),
					Topics: project.Topics,
				})
			}

			return projects, err
		},
	},
	"bgg": {
		Request: Request{
			Method: "GET",
			Path:   "https://boardgamegeek.com/xmlapi/geeklist/332832",
		},
		Parser: func(data []byte) (projects []Project, err error) {
			var projectItems []struct {
				Item struct {
					Id string `xml:"objectid,attr"`
				} `xml:"item"`
			}
			if unmarshalErr := xml.Unmarshal(data, &projectItems); unmarshalErr != nil {
				return nil, errors.Join(errors.New("error parsing BGG projects"), unmarshalErr)
			}

			for _, item := range projectItems {
				projectData, _ := Fetch(&Request{
					Method: "GET",
					Path:   fmt.Sprintf("https://boardgamegeek.com/xmlapi/boardgame/%s", item.Item.Id),
				})
				var bggProject struct {
					Title    string   `xml:"boardgame>name"`
					ImageSrc string   `xml:"boardgame>image"`
					Tags     []string `xml:"boardgame>boardgamemechanic"`
				}

				if unmarshalErr := xml.Unmarshal(projectData, &bggProject); unmarshalErr != nil {
					err = errors.Join(err, fmt.Errorf("error parsing BGG project (%s)", item.Item.Id), unmarshalErr)
					continue
				}

				projects = append(projects, Project{
					Host:  "Board Game Geek",
					Title: bggProject.Title,
					Url:   fmt.Sprintf("https://boardgamegeek.com/boardgame/%s", item.Item.Id),
					Image: Image{
						Src: template.HTMLAttr(fmt.Sprintf("src=\"%s\"", bggProject.ImageSrc)),
						Alt: template.HTMLAttr(fmt.Sprintf("alt=\"Board Game: %s\"", bggProject.Title)),
					},
					Logo:   GetSVGLogo("bgg"),
					Topics: bggProject.Tags,
				})
			}

			return projects, err
		},
	},
}
