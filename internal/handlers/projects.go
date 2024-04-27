package handlers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	githublangsgo "github.com/NDoolan360/github-langs-go"
	"github.com/NDoolan360/go-htmx-site/internal/components"
	"github.com/a-h/templ"
)

// Projects handles the request for fetching and rendering project data.
func Projects(w http.ResponseWriter, r *http.Request) {
	projects, errs := FetchProjects(r.URL.Query()["host"])
	// TODO render projects and send to client with sse
	if len(projects) > 0 {
		for _, project := range projects {
			components.ProjectTemplate(project).Render(r.Context(), w)
		}
	} else if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
	}
}

type Host struct {
	Request
	Parser func([]byte) ([]components.Project, error)
}

type Request struct {
	Method      string
	Path        string
	Username    string
	Password    string
	ContentType string
	Body        string
}

func FetchProjects(hostNames []string) (projects []components.Project, errs error) {
	for _, hostName := range hostNames {
		host, ok := hostMap[hostName]
		if ok {
			data, err := Fetch(&host.Request)
			if err != nil {
				errs = errors.Join(errs, err)
				continue
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
		outgoingRequest.Header.Add("Content-Type", request.ContentType)
	}
	if len(request.Username) > 0 && len(request.Password) > 0 {
		outgoingRequest.SetBasicAuth(request.Username, request.Password)
	}

	client := &http.Client{}
	response, err := client.Do(outgoingRequest)
	if err != nil {
		return nil, err
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request to %s failed with status code: %d", request.Path, response.StatusCode)
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
		Parser: func(data []byte) (projects []components.Project, err error) {
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

				lang, languageErr := githublangsgo.GetLanguage(project.Language)

				if languageErr != nil {
					err = errors.Join(err, fmt.Errorf("error parsing language (%s)", project.Language), languageErr)
				}

				projects = append(projects, components.Project{
					Host:        "Github",
					Title:       project.Title,
					Description: project.Description,
					Url:         templ.SafeURL(project.Url),
					Language: components.Language{
						Name:   project.Language,
						Colour: lang.Color,
					},
					Logo:   components.GithubLogo(),
					Topics: project.Topics,
				})
			}

			return projects, err
		},
	},
	"cults3d": {
		Request: Request{
			Method:      "POST",
			Path:        "https://cults3d.com/graphql",
			Username:    os.Getenv("CULTS3D_USERNAME"),
			Password:    os.Getenv("CULTS3D_API_KEY"),
			Body:        "{\"query\":\"{user(nick:\\\"ND360\\\"){creations(limit:5,sort:BY_DOWNLOADS){name url illustrationImageUrl tags}}}\"}",
			ContentType: "application/json",
		},
		Parser: func(data []byte) (projects []components.Project, err error) {
			var cults3dProjects struct {
				Data struct {
					User struct {
						Creations []struct {
							Title       string   `json:"name"`
							Description string   `json:"description"`
							Url         string   `json:"url"`
							ImageSrc    string   `json:"illustrationImageUrl"`
							Topics      []string `json:"tags"`
						} `json:"creations"`
					} `json:"user"`
				} `json:"data"`
			}
			if unmarshalErr := json.Unmarshal(data, &cults3dProjects); unmarshalErr != nil {
				return nil, errors.Join(errors.New("error parsing Cults3D projects"), unmarshalErr)
			}

			for _, project := range cults3dProjects.Data.User.Creations {
				projects = append(projects, components.Project{
					Host:        "Cults3D",
					Title:       project.Title,
					Description: project.Description,
					Url:         templ.SafeURL(project.Url),
					Image: components.Image{
						Src: project.ImageSrc,
						Alt: fmt.Sprintf("3D Model: %s", project.Title),
					},
					Logo:   components.Cults3DLogo(),
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
		Parser: func(data []byte) (projects []components.Project, err error) {
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

				projects = append(projects, components.Project{
					Host:  "Board Game Geek",
					Title: bggProject.Title,
					Url:   templ.SafeURL(fmt.Sprintf("https://boardgamegeek.com/boardgame/%s", item.Item.Id)),
					Image: components.Image{
						Src: bggProject.ImageSrc,
						Alt: fmt.Sprintf("Board Game: %s", bggProject.Title),
					},
					Logo:   components.BGGLogo(),
					Topics: bggProject.Tags,
				})
			}

			return projects, err
		},
	},
}
