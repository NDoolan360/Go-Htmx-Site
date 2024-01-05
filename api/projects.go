package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Project struct {
	Host  string
	Logo  string
	Image struct {
		Href string
		Alt  string
	}
	Title          string   `json:"name"`
	Description    string   `json:"description"`
	HtmlUrl        string   `json:"html_url"`
	Topics         []string `json:"topics"`
	Language       string   `json:"language"`
	LanguageColour string
}

type UrlDetails struct {
	Name string
	Path string
	Type string
}

var HostMap = map[string]UrlDetails{
	"github": {
		Name: "Github",
		Path: "https://api.github.com/users/NDoolan360/repos?sort=stars",
		Type: "json"},
	"cults3d": {
		Name: "Cults3d",
		Path: "https://cults3d.com/en/users/ND360/3d-models",
		Type: "html"},
	"bgg": {
		Name: "Board Game Geek",
		Path: "https://boardgamegeek.com/geeksearch.php?action=search&advsearch=1&objecttype=boardgame&include%5Bdesignerid%5D=133893",
		Type: "html",
	},
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	hosts := r.URL.Query()["host"]
	for _, host := range hosts {
		if site, ok := HostMap[host]; !ok {
			http.Error(w, fmt.Sprintf("URL not found for host: %s", host), http.StatusNotFound)
		} else if content, err := Fetch(site); err != nil {
			http.Error(w, fmt.Sprintf("error fetching content from host %s: %s", host, err.Error()), http.StatusInternalServerError)
		} else if projects, err := Parse(content, host); err != nil {
			http.Error(w, fmt.Sprintf("error parsing content from host %s: %s", host, err.Error()), http.StatusInternalServerError)
		} else {
			for _, project := range projects {
				// TODO use template to return html
				project.Host = site.Name
				project.Logo = fmt.Sprintf("/images/logos/%s.svg", host)
				if project.Language != "" {
					project.LanguageColour = "Colour"
				}
				fmt.Fprintf(w, "%s Project:\n%+v\n\n", project.Host, *project)
			}
		}
	}
}

func Fetch(url UrlDetails) (string, error) {
	if resp, err := http.Get(url.Path); err != nil {
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

func Parse(content string, host string) ([]*Project, error) {
	var projects []*Project
	var err error

	switch host {
	case "github":
		err = json.Unmarshal([]byte(content), &projects)
	default:
		err = fmt.Errorf("unsupported host")
	}

	if err != nil {
		return nil, err
	}
	return projects, nil
}
