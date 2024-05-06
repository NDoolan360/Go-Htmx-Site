package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	githublangsgo "github.com/NDoolan360/github-langs-go"
	"github.com/NDoolan360/go-htmx-site/website/components"
	"github.com/a-h/templ"
)

type GithubHost struct {
	BaseURL string
	User    string
}

func (gh GithubHost) Fetch() ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s/repos?sort=stars", gh.BaseURL, gh.User))
	if err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

func (_ GithubHost) Parse(data []byte) (projects []Project, err error) {
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

		projects = append(projects, Project{
			Host:        "Github",
			Title:       project.Title,
			Description: project.Description,
			Url:         templ.SafeURL(project.Url),
			Language: Language{
				Name:   project.Language,
				Colour: lang.Color,
			},
			Logo:   components.GithubLogo(),
			Topics: project.Topics,
		})
	}

	return projects, err
}
