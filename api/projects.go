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

	githublangsgo "github.com/NDoolan360/github-langs-go"
	"golang.org/x/net/html"
)

// ProjectsTemplate represents the data structure for the projects.gohtml template.
type ProjectsTemplate struct {
	Projects []Project
}

// Project represents information about a personal project.
type Project struct {
	Host           string
	LogoSVG        template.HTML
	ImageSrc       template.HTMLAttr
	ImageAlt       template.HTMLAttr
	Title          string   `json:"name"`
	Description    string   `json:"description"`
	HtmlUrl        string   `json:"html_url"`
	Topics         []string `json:"topics"`
	Fork           bool     `json:"fork"`
	Language       string   `json:"language"`
	LanguageColour template.CSS
}

// Date represents a start and end date for an experience.
type Date struct {
	Start string
	End   string
}

// Position represents a job title and current job status for work experience.
type Position struct {
	Title   string
	Current bool
}

// GetProjects handles the request for fetching and rendering project data.
func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, errs := FetchProjects(r.URL.Query()["host"])
	if len(errs) > 0 {
		var errorMessages string
		for _, err := range errs {
			errorMessages += err.Error() + "\n"
		}
		http.Error(w, errorMessages, http.StatusInternalServerError)
	} else {
		tmpl := template.Must(template.ParseFiles(
			GetApiAsset("template/projects.gohtml"),
		))
		err := tmpl.Execute(w, ProjectsTemplate{Projects: projects})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// HostMap contains information about different project hosts for fetching content.
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

func FetchProjects(hosts []string) ([]Project, []error) {
	projects := []Project{}
	errs := []error{}
	for _, host := range hosts {
		if site, ok := HostMap[host]; !ok {
			errs = append(errs, fmt.Errorf("URL not found for host: %s", host))

		} else if content, fetchErr := Fetch(site.Path); fetchErr != nil {
			errs = append(errs, fmt.Errorf("error fetching content from host %s: %s", host, fetchErr.Error()))

		} else if hostProjects, ParseErr := Parse(content, host, site.Type); ParseErr != nil {
			errs = append(errs, fmt.Errorf("error parsing content from host %s: %s", host, ParseErr.Error()))

		} else {
			for _, project := range hostProjects {
				// Skip unimportant Github Repos
				if host == "github" && (project.Fork || len(project.Topics) == 0) {
					continue
				}
				project.Host = site.Name
				project.LogoSVG = GetSVGLogo(host)
				if project.Language != "" {
					if lang, err := githublangsgo.GetLanguage(project.Language); err == nil {
						project.LanguageColour = template.CSS(lang.Color)
					}
				}
				if host == "bgg" {
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
			projects = ParseHTMLDoc(doc, BGGNode)
		case "cults3d":
			projects = ParseHTMLDoc(doc, Cults3DNode)
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

type BoardGame struct {
	ObjectID  string   `xml:"objectid,attr"`
	Mechanics []string `xml:"boardgamemechanic"`
	ImageURL  string   `xml:"image"`
}

type BoardGames struct {
	TermsOfUse string    `xml:"termsofuse,attr"`
	BoardGame  BoardGame `xml:"boardgame"`
}

func UpgradeBGG(project *Project) error {
	re := regexp.MustCompile(`/boardgame/(\d+)`)
	matches := re.FindStringSubmatch(project.HtmlUrl)
	if len(matches) < 2 {
		return fmt.Errorf("Boardgame ID not found in the URL")
	}
	boardgameID := matches[1]

	bggXML, fetchErr := Fetch(fmt.Sprintf("https://api.geekdo.com/xmlapi/boardgame/%s", boardgameID))
	if fetchErr != nil {
		return fetchErr
	}

	var boardGames BoardGames
	parseErr := xml.Unmarshal([]byte(bggXML), &boardGames)
	if parseErr != nil {
		return fmt.Errorf("error unmarshaling XML: %s", parseErr)
	}

	if boardGames.BoardGame.ImageURL != "" {
		project.ImageSrc = template.HTMLAttr(fmt.Sprintf(`src="%s"`, boardGames.BoardGame.ImageURL))
	}
	if len(boardGames.BoardGame.Mechanics) > 0 {
		project.Topics = boardGames.BoardGame.Mechanics
	}

	return nil
}
