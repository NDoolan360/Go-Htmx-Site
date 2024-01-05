package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	utils "github.com/ndoolan360/ndsite/src"
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
	if projects, err := FetchAllProjects(r.URL.Query()["host"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		projectJSON, _ := json.MarshalIndent(projects, "", "  ")
		fmt.Fprintf(w, "%s\n\n", string(projectJSON))
	}
}

func FetchAllProjects(hosts []string) ([]*Project, error) {
	var projects []*Project
	for _, host := range hosts {
		if site, ok := HostMap[host]; !ok {
			return nil, fmt.Errorf("URL not found for host: %s", host)
		} else if content, err := utils.Fetch(site.Path); err != nil {
			return nil, fmt.Errorf("error fetching content from host %s: %s", host, err.Error())
		} else if hostProjects, err := Parse(content, host); err != nil {
			return nil, fmt.Errorf("error parsing content from host %s: %s", host, err.Error())
		} else {
			for _, project := range hostProjects {
				// TODO use template to return html
				project.Host = site.Name
				project.Logo = fmt.Sprintf("/images/logos/%s.svg", host)
				if project.Language != "" {
					// TODO Map the Language name to a LanguageColour
					project.LanguageColour = "Colour"
				}
				// Skip unimportant Github Repos
				if host == "github" && (project.Fork || len(project.Topics) == 0) {
					continue
				}
				projects = append(projects, project)
			}
		}
	}
	return projects, nil
}

func Parse(content string, host string) ([]*Project, error) {
	var projects []*Project
	var err error

	switch host {
	case "github":
		err = json.Unmarshal([]byte(content), &projects)
	case "bgg", "cults3d":
		doc, parseErr := html.Parse(strings.NewReader(content))
		if parseErr != nil {
			err = fmt.Errorf("error parsing HTML: %s", parseErr)
		}
		switch host {
		case "bgg":
			projects = utils.ParseHTMLDoc[Project](doc, BGGNode)
		case "cults3d":
			projects = utils.ParseHTMLDoc[Project](doc, Cults3DNode)
		}
	default:
		err = fmt.Errorf("unsupported host")
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
