package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"

	githublangsgo "github.com/NDoolan360/github-langs-go"
	"golang.org/x/net/html"
)

type Projects struct {
	Projects []Project
}

type Project struct {
	Host    string
	LogoSVG template.HTML
	Image   struct {
		Src     template.HTMLAttr
		AltText string
	}
	Title          string   `json:"name"`
	Description    string   `json:"description"`
	HtmlUrl        string   `json:"html_url"`
	Topics         []string `json:"topics"`
	Fork           bool     `json:"fork"`
	Language       string   `json:"language"`
	LanguageColour template.CSS
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	hosts := r.URL.Query()["host"]
	projects, errs := FetchAllProjects(hosts)
	if errs != nil {
		var errorMessages string
		for _, err := range errs {
			errorMessages += err.Error() + "\n"
		}
		http.Error(w, errorMessages, http.StatusInternalServerError)
	} else {
		tmp := template.Must(template.ParseFiles(
			FromPWD("api/template/projects.gohtml"),
		))
		err := tmp.Execute(w, Projects{projects})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func FetchAllProjects(hosts []string) (projects []Project, err []error) {
	hostMap := map[string]struct {
		Name string
		Path string
		Type string
	}{
		"github": {
			Name: "Github",
			Path: "https://api.github.com/users/NDoolan360/repos?sort=stars",
			Type: "json",
		},
		"bgg": {
			Name: "Board Game Geek",
			Path: "https://boardgamegeek.com/geeksearch.php?action=search&advsearch=1&objecttype=boardgame&include%5Bdesignerid%5D=133893",
			Type: "html",
		},
		"cults3d": {
			Name: "Cults3D",
			Path: "https://cults3d.com/en/users/ND360/3d-models",
			Type: "html",
		},
	}

	for _, host := range hosts {
		if site, ok := hostMap[host]; !ok {
			err = append(err, fmt.Errorf("URL not found for host: %s", host))
		} else if content, fetchErr := Fetch(site.Path); err != nil {
			err = append(err, fmt.Errorf("error fetching content from host %s: %s", host, fetchErr.Error()))
		} else if hostProjects, ParseErr := Parse(content, host, site.Type); err != nil {
			err = append(err, fmt.Errorf("error parsing content from host %s: %s", host, ParseErr.Error()))
		} else {
			for _, project := range hostProjects {
				// Skip unimportant Github Repos
				if host == "github" && (project.Fork || len(project.Topics) == 0) {
					continue
				}
				project.Host = site.Name
				if svg, err := GetSVGLogo(host); err == nil {
					project.LogoSVG = svg
				}
				if project.Language != "" {
					if lang, err := githublangsgo.GetLanguage(project.Language); err == nil {
						project.LanguageColour = template.CSS(lang.Color)
					}
				}
				projects = append(projects, *project)
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
			projects = ParseHTMLDoc[Project](doc, BGGNode)
		case "cults3d":
			projects = ParseHTMLDoc[Project](doc, Cults3DNode)
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
	if node.Data == "tr" && WithTagEqual("id", "row_")(node) {
		project := Project{}
		if title := FirstInChildren(node, WithTagEqual("class", "primary")); title != nil {
			project.Title = GetTextContent(title)
		}
		if description := FirstInChildren(node, WithTagEqual("class", "smallefont")); description != nil {
			project.Description = GetTextContent(description)
		}
		if thumbnail := FirstInChildren(node, WithTagEqual("class", "collection_thumbnail")); thumbnail != nil {
			if a := FirstInChildren(thumbnail, WithTag("a")); a != nil {
				project.HtmlUrl = fmt.Sprintf("https://boardgamegeek.com%s", GetAttribute(a, "href"))
				if img := FirstInChildren(a, WithTag("img")); img != nil {
					project.Image.Src = template.HTMLAttr(fmt.Sprintf("src=\"%s\"", GetAttribute(img, "src")))
					project.Image.AltText = GetAttribute(img, "alt")
				}
			}
		}
		return &project, true
	}

	return nil, false
}

func Cults3DNode(node *html.Node) (*Project, bool) {
	if node.Data == "article" && WithTagEqual("class", "crea")(node) {
		project := Project{}
		if h3 := FirstInChildren(node, WithTag("h3")); h3 != nil {
			project.Title = GetTextContent(h3)
		}
		if a := FirstInChildren(node, WithTag("a")); a != nil {
			project.HtmlUrl = fmt.Sprintf("https://cults3d.com%s", GetAttribute(a, "href"))
		}
		if img := FirstInChildren(node, WithTag("img")); img != nil {
			dataSrc := GetAttribute(img, "data-src")

			// extract full size file rather than thumbnail image if possible
			regex := regexp.MustCompile(`https://files\.cults3d\.com[^'"]+`)
			match := regex.FindString(dataSrc)

			if match != "" {
				project.Image.Src = template.HTMLAttr(fmt.Sprintf("src=\"%s\"", match))
			} else {
				project.Image.Src = template.HTMLAttr(fmt.Sprintf("src=\"%s\"", dataSrc))
			}

			project.Image.AltText = GetAttribute(img, "alt")
		}
		return &project, true
	}

	return nil, false
}
