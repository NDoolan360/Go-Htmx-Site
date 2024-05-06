package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/NDoolan360/go-htmx-site/website/components"
	"github.com/a-h/templ"
)

type Cults3dHost struct {
	BaseURL string
	User    string
}

func (cults Cults3dHost) Fetch() ([]byte, error) {
	client := &http.Client{}
	body := fmt.Sprintf("{\"query\":\"{user(nick:\\\"%s\\\"){creations(limit:5,sort:BY_DOWNLOADS){name url illustrationImageUrl tags}}}\"}", cults.User)
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/graphql", cults.BaseURL), strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(os.Getenv("CULTS3D_USERNAME"), os.Getenv("CULTS3D_API_KEY"))

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request to cults3d failed with status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	if data, err := io.ReadAll(response.Body); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (_ Cults3dHost) Parse(data []byte) (projects []Project, err error) {
	var cults3dProjects struct {
		Data struct {
			User struct {
				Creations []struct {
					Title       string   `json:"name"`
					Description string   `json:"description"`
					Url         string   `json:"url"`
					ImageSrc    string   `json:"illustrationImageUrl"`
					Topics      []string `json:"tags"`
				} `json:"creations"`
			} `json:"user"`
		} `json:"data"`
	}
	if unmarshalErr := json.Unmarshal(data, &cults3dProjects); unmarshalErr != nil {
		return nil, errors.Join(errors.New("error parsing Cults3D projects"), unmarshalErr)
	}

	for _, project := range cults3dProjects.Data.User.Creations {
		projects = append(projects, Project{
			Host:        "Cults3D",
			Title:       project.Title,
			Description: project.Description,
			Url:         templ.URL(project.Url),
			Image: Image{
				Src: project.ImageSrc,
				Alt: fmt.Sprintf("3D Model: %s", project.Title),
			},
			Logo:   components.Logo("Cults3d"),
			Topics: project.Topics,
		})
	}

	return projects, err
}
