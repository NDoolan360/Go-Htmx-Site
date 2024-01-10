package utils

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

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

// content fetching

func FetchURL(url string) (string, error) {
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

// content parsing

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

// html parsing

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
					project.ImageAttr = []template.HTMLAttr{
						template.HTMLAttr(fmt.Sprintf("src=\"%s\"", GetAttribute(img, "src"))),
						template.HTMLAttr(fmt.Sprintf("alt=\"%s\"", GetAttribute(img, "alt"))),
					}
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
			if match := regex.FindString(dataSrc); match != "" {
				dataSrc = match
			}

			project.ImageAttr = []template.HTMLAttr{
				template.HTMLAttr(fmt.Sprintf("src=\"%s\"", dataSrc)),
				template.HTMLAttr(fmt.Sprintf("alt=\"%s\"", GetAttribute(img, "alt"))),
			}
		}
		return &project, true
	}

	return nil, false
}

// html helpers

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
