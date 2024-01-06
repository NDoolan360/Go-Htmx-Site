package src

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

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
