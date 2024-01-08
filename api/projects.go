package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"

	githublangsgo "github.com/NDoolan360/github-langs-go"
	utils "github.com/NDoolan360/go-htmx-site/src"
	"golang.org/x/net/html"
)

type Project struct {
	Host  string
	Logo  string
	Image struct {
		Src string
		Alt string
	}
	Title          string   `json:"name"`
	Description    string   `json:"description"`
	HtmlUrl        string   `json:"html_url"`
	Topics         []string `json:"topics"`
	Fork           bool
	Language       string `json:"language"`
	LanguageColour string
}

var HostMap = map[string]struct {
	Name string
	Path string
	Type string
}{
	"github": {
		Name: "Github",
		Path: "https://api.github.com/users/NDoolan360/repos?sort=stars",
		Type: "json",
	},
	"cults3d": {
		Name: "Cults3d",
		Path: "https://cults3d.com/en/users/ND360/3d-models",
		Type: "html",
	},
	"bgg": {
		Name: "Board Game Geek",
		Path: "https://boardgamegeek.com/geeksearch.php?action=search&advsearch=1&objecttype=boardgame&include%5Bdesignerid%5D=133893",
		Type: "html",
	},
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	hosts := r.URL.Query()["host"]
	if len(hosts) == 0 {
		fmt.Fprint(w, "No hosts specified in query params.")
	} else if projects, errs := FetchAllProjects(hosts); projects == nil && errs != nil {
		var errorMessages string
		for _, err := range errs {
			errorMessages += err.Error() + "\n"
		}
		http.Error(w, errorMessages, http.StatusInternalServerError)
	} else if projects != nil {
		if tmpl, parseErr := template.ParseFiles("template/project.gohtml"); parseErr == nil {
			for _, project := range projects {
				tmpl.Execute(w, project)
			}
		}
	} else {
		fmt.Fprint(w, "No projects found.")
	}
}

func FetchAllProjects(hosts []string) (projects []*Project, err []error) {
	for _, host := range hosts {
		if site, ok := HostMap[host]; !ok {
			err = append(err, fmt.Errorf("URL not found for host: %s", host))
		} else if content, fetchErr := utils.Fetch(site.Path); err != nil {
			err = append(err, fmt.Errorf("error fetching content from host %s: %s", host, fetchErr.Error()))
		} else if hostProjects, ParseErr := Parse(content, host, site.Type); err != nil {
			err = append(err, fmt.Errorf("error parsing content from host %s: %s", host, ParseErr.Error()))
		} else {
			for _, project := range hostProjects {
				project.Host = site.Name
				project.Logo = fmt.Sprintf("/images/logos/%s.svg", host)
				if project.Language != "" {
					if lang, err := githublangsgo.GetLanguage(project.Language); err == nil {
						project.LanguageColour = lang.Color
					}
				}
				// Skip unimportant Github Repos
				if host == "github" && (project.Fork || len(project.Topics) == 0) {
					continue
				}
				projects = append(projects, project)
			}
		}
	}
	return projects, err
}

func Parse(content string, host string, contentType string) ([]*Project, error) {
	var projects []*Project
	var err error

	switch contentType {
	case "json":
		err = json.Unmarshal([]byte(content), &projects)
	case "html":
		doc, parseErr := html.Parse(strings.NewReader(content))
		if parseErr != nil {
			err = fmt.Errorf("error parsing HTML: %s", parseErr)
		}
		switch host {
		case "bgg":
			projects = utils.ParseHTMLDoc[Project](doc, BGGNode)
		case "cults3d":
			projects = utils.ParseHTMLDoc[Project](doc, Cults3DNode)
		default:
			err = fmt.Errorf("unsupported host")
		}
	default:
		err = fmt.Errorf("unsupported content type: %s", contentType)
	}

	if err != nil {
		return nil, err
	}
	return projects, nil
}

func BGGNode(node *html.Node) (*Project, bool) {
	if node.Data == "tr" && strings.Contains(utils.GetAttribute(node, "id"), "row_") {
		project := Project{}
		if title := utils.FirstInChildren(node, utils.WithClass("primary")); title != nil {
			project.Title = utils.GetTextContent(title)
		}
		if description := utils.FirstInChildren(node, utils.WithClass("smallefont")); description != nil {
			project.Description = utils.GetTextContent(description)
		}
		if thumbnail := utils.FirstInChildren(node, utils.WithClass("collection_thumbnail")); thumbnail != nil {
			if link := thumbnail.FirstChild; link != nil {
				project.HtmlUrl = "https://boardgamegeek.com" + utils.GetAttribute(link, "href")
				if img := link.FirstChild; img != nil {
					project.Image.Src = utils.GetAttribute(img, "src")
					project.Image.Alt = utils.GetAttribute(img, "alt")
				}
			}
		}
		return &project, true
	}

	return nil, false
}

func Cults3DNode(node *html.Node) (*Project, bool) {
	if node.Data == "article" && strings.Contains(utils.GetAttribute(node, "class"), "crea") {
		project := Project{}
		if h3 := utils.FirstInChildren(node, utils.WithTag("h3")); h3 != nil {
			project.Title = utils.GetTextContent(h3)
		}
		if a := utils.FirstInChildren(node, utils.WithTag("a")); a != nil {
			project.HtmlUrl = "https://cults3d.com" + utils.GetAttribute(a, "href")
		}
		if img := utils.FirstInChildren(node, utils.WithTag("img")); img != nil {
			project.Image.Src = utils.GetAttribute(img, "data-src")

			// extract full size file rather than thumbnail image if possible
			regex := regexp.MustCompile(`https://files\.cults3d\.com[^'"]+`)
			match := regex.FindString(project.Image.Src)

			if match != "" {
				project.Image.Src = match
			}

			project.Image.Alt = utils.GetAttribute(img, "alt")
		}
		return &project, true
	}

	return nil, false
}
