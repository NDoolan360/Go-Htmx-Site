package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strings"

	githubLanguage "github.com/NDoolan360/github-langs-go"
	"golang.org/x/net/html"
)

// GetProjects handles the request for fetching and rendering project data.
func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, errs := FetchProjects(r.URL.Query()["host"])
	if len(projects) > 0 {
		projectTemplate := template.Must(template.ParseFiles(
			GetApiAsset("template/projects.gohtml"),
		))
		err := projectTemplate.Execute(w, ProjectsTemplate{Projects: projects})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		errorMessages := ""
		for _, err := range errs {
			errorMessages += err.Error() + "\n"
		}
		http.Error(w, errorMessages, http.StatusInternalServerError)
	}
}

// HostMap contains information about different project hosts for fetching and parsing content.
var HostMap = map[string]Host{
	"github": {
		Name: "Github",
		Path: "https://api.github.com/users/NDoolan360/repos?sort=stars",
		Type: "json",
	},
	"bgg": {
		Name:   "Board Game Geek",
		Path:   "https://boardgamegeek.com/geeksearch.php?action=search&advsearch=1&objecttype=boardgame&include%5Bdesignerid%5D=133893",
		Type:   "html",
		Parser: BGGNode,
	},
	"cults3d": {
		Name:   "Cults3D",
		Path:   "https://cults3d.com/en/users/ND360/3d-models",
		Type:   "html",
		Parser: Cults3DNode,
	},
}

func FetchProjects(hostNames []string) ([]Project, []error) {
	projects := []Project{}
	errs := []error{}
	for _, hostName := range hostNames {
		if host, ok := HostMap[hostName]; !ok {
			errs = append(errs, fmt.Errorf("URL not found for host: %s", hostName))

		} else if content, fetchErr := Fetch(host.Path); fetchErr != nil {
			errs = append(errs, fmt.Errorf("error fetching content from host %s: %s", hostName, fetchErr.Error()))

		} else if hostProjects, ParseErr := Parse(host, content); ParseErr != nil {
			errs = append(errs, fmt.Errorf("error parsing content from host %s: %s", hostName, ParseErr.Error()))

		} else {
			for _, project := range hostProjects {
				// Skip unimportant Github Repos
				if hostName == "github" && (project.Fork || len(project.Topics) == 0) {
					continue
				}
				project.Host = host.Name
				project.LogoSVG = GetSVGLogo(hostName)
				if project.Language != "" {
					if lang, err := githubLanguage.GetLanguage(project.Language); err == nil {
						project.LanguageColour = template.CSS(lang.Color)
					}
				}
				if hostName == "bgg" {
					if err := UpgradeBGG(project); err != nil {
						continue
					}
				}
				projects = append(projects, *project)
			}
		}
	}

	return projects, errs
}

// Fetch fetches content from a given URL using HTTP GET.
var Fetch = func(url string) (string, error) {
	if resp, err := http.Get(url); err != nil {
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

// Parse parses content based on host and content type, returning a list of projects.
func Parse(host Host, content string) ([]*Project, error) {
	var projects []*Project
	var err error
	switch host.Type {
	case "json":
		err = json.Unmarshal([]byte(content), &projects)
	case "html":
		doc, parseErr := html.Parse(strings.NewReader(content))
		if parseErr != nil {
			err = fmt.Errorf("error parsing HTML: %s", parseErr)
		}
		if host.Parser != nil {
			projects = ParseHTMLDoc(doc, host.Parser)
		} else {
			err = fmt.Errorf("no parser provided for host: %s", host.Name)
		}
	default:
		err = fmt.Errorf("unsupported content type: %s", host.Type)
	}

	if err != nil {
		return nil, err
	}
	return projects, nil
}

// Functions for HTML parsing and helpers.

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

// BGGNode is a parser function for Board Game Geek (BGG) HTML content.
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
					project.ImageSrc = template.HTMLAttr(fmt.Sprintf("src=\"%s\"", GetAttribute(img, "src")))
					project.ImageAlt = template.HTMLAttr(fmt.Sprintf("alt=\"%s\"", GetAttribute(img, "alt")))
				}
			}
		}
		return &project, true
	}

	return nil, false
}

// Cults3DNode is a parser function for Cults3D HTML content.
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
			if match := regex.FindString(dataSrc); match != "" {
				dataSrc = match
			}

			project.ImageSrc = template.HTMLAttr(fmt.Sprintf("src=\"%s\"", dataSrc))
			project.ImageAlt = template.HTMLAttr(fmt.Sprintf("alt=\"%s\"", GetAttribute(img, "alt")))
		}
		return &project, true
	}

	return nil, false
}

// HTML parsing helpers.

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

func WithTagEqual(tag string, value string) MatchPredicate {
	return func(node *html.Node) bool {
		return strings.Contains(GetAttribute(node, tag), value)
	}
}

func UpgradeBGG(project *Project) error {
	re := regexp.MustCompile(`/boardgame/(\d+)`)
	matches := re.FindStringSubmatch(project.HtmlUrl)
	if len(matches) < 2 {
		return fmt.Errorf("boardgame ID not found in the URL")
	}
	boardgameID := matches[1]

	bggXML, fetchErr := Fetch(fmt.Sprintf("https://api.geekdo.com/xmlapi/boardgame/%s", boardgameID))
	if fetchErr != nil {
		return fetchErr
	}

	var boardGames BoardGames
	parseErr := xml.Unmarshal([]byte(bggXML), &boardGames)
	if parseErr != nil {
		return fmt.Errorf("error un-marshaling XML: %s", parseErr)
	}

	if boardGames.BoardGame.ImageURL != "" {
		project.ImageSrc = template.HTMLAttr(fmt.Sprintf(`src="%s"`, boardGames.BoardGame.ImageURL))
	}
	if len(boardGames.BoardGame.Mechanics) > 0 {
		project.Topics = boardGames.BoardGame.Mechanics
	}

	return nil
}
