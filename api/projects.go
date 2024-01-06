package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

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
		projectJSON, _ := json.MarshalIndent(projects, "", "  ")
		fmt.Fprintf(w, "%s\n\n", string(projectJSON))
	} else {
		fmt.Fprint(w, "No projects found.")
	}
}

func FetchAllProjects(hosts []string) (projects []*Project, err []error) {
	for _, host := range hosts {
		if site, ok := HostMap[host]; !ok {
			err = append(err, fmt.Errorf("URL not found for host: %s", host))
		} else if content, fetchErr := Fetch(site.Path); err != nil {
			err = append(err, fmt.Errorf("error fetching content from host %s: %s", host, fetchErr.Error()))
		} else if hostProjects, ParseErr := Parse(content, host, site.Type); err != nil {
			err = append(err, fmt.Errorf("error parsing content from host %s: %s", host, ParseErr.Error()))
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
	return projects, err
}

func Fetch(path string) (string, error) {
	if resp, err := http.Get(path); err != nil {
		return "", err
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
		} else if body, err := io.ReadAll(resp.Body); err != nil {
			return "", err
		} else {
			return string(body), nil
		}
	}
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

type NodePredicate[T any] func(*html.Node) (*T, bool)

func ParseHTMLDoc[T any](node *html.Node, check NodePredicate[T]) (collection []*T) {
	if object, ok := check(node); ok {
		collection = append(collection, object)
	}

	for next := node.FirstChild; next != nil; next = next.NextSibling {
		collection = append(collection, ParseHTMLDoc(next, check)...)
	}

	return collection
}

func BGGNode(node *html.Node) (*Project, bool) {
	if node.Data == "tr" && strings.Contains(GetAttribute(node, "id"), "row_") {
		project := Project{}
		if title := FirstInChildren(node, WithClass("primary")); title != nil {
			project.Title = GetTextContent(title)
		}
		if description := FirstInChildren(node, WithClass("smallefont")); description != nil {
			project.Description = GetTextContent(description)
		}
		if thumbnail := FirstInChildren(node, WithClass("collection_thumbnail")); thumbnail != nil {
			if link := thumbnail.FirstChild; link != nil {
				project.HtmlUrl = "https://boardgamegeek.com" + GetAttribute(link, "href")
				if img := link.FirstChild; img != nil {
					project.Image.Src = GetAttribute(img, "src")
					project.Image.Alt = GetAttribute(img, "alt")
				}
			}
		}
		return &project, true
	}

	return nil, false
}

func Cults3DNode(node *html.Node) (*Project, bool) {
	if node.Data == "article" && strings.Contains(GetAttribute(node, "class"), "crea") {
		project := Project{}
		if h3 := FirstInChildren(node, WithTag("h3")); h3 != nil {
			project.Title = GetTextContent(h3)
		}
		if a := FirstInChildren(node, WithTag("a")); a != nil {
			project.HtmlUrl = "https://cults3d.com" + GetAttribute(a, "href")
		}
		if img := FirstInChildren(node, WithTag("img")); img != nil {
			project.Image.Src = GetAttribute(img, "data-src")

			// extract full size file rather than thumbnail image if possible
			regex := regexp.MustCompile(`https://files\.cults3d\.com[^'"]+`)
			match := regex.FindString(project.Image.Src)

			if match != "" {
				project.Image.Src = match
			}

			project.Image.Alt = GetAttribute(img, "alt")
		}
		return &project, true
	}

	return nil, false
}

func GetAttribute(n *html.Node, attrName string) string {
	for _, attr := range n.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}

func GetTextContent(n *html.Node) string {
	return strings.TrimSpace(n.FirstChild.Data)
}

type MatchPredicate func(*html.Node) bool

func FirstInChildren(node *html.Node, match MatchPredicate) *html.Node {
	if node == nil {
		return nil
	}
	if match(node) {
		return node
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if found := FirstInChildren(c, match); found != nil {
			return found
		}
	}

	return nil
}

func WithTag(tag string) MatchPredicate {
	return func(node *html.Node) bool {
		return node.Data == tag
	}
}

func WithClass(class string) MatchPredicate {
	return func(node *html.Node) bool {
		return strings.Contains(GetAttribute(node, "class"), class)
	}
}
