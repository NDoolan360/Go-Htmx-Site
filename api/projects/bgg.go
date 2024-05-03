package main

import (
	_ "embed"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/website/components"
	"github.com/a-h/templ"
)

type BggHost struct {
	BaseURL  string
	Geeklist string
}

func (bgg BggHost) Fetch() ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/geeklist/%s", bgg.BaseURL, bgg.Geeklist))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (bgg BggHost) Parse(data []byte) (projects []Project, err error) {
	var projectItems []struct {
		Item struct {
			Id string `xml:"objectid,attr"`
		} `xml:"item"`
	}
	if unmarshalErr := xml.Unmarshal(data, &projectItems); unmarshalErr != nil {
		return nil, errors.Join(errors.New("error parsing BGG projects"), unmarshalErr)
	}

	for _, item := range projectItems {
		resp, err := http.Get(fmt.Sprintf("%s/boardgame/%s", bgg.BaseURL, item.Item.Id))
		if err != nil {
			return nil, err
		}

		projectData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var bggProject struct {
			Title    string   `xml:"boardgame>name"`
			ImageSrc string   `xml:"boardgame>image"`
			Tags     []string `xml:"boardgame>boardgamemechanic"`
		}

		if unmarshalErr := xml.Unmarshal(projectData, &bggProject); unmarshalErr != nil {
			err = errors.Join(err, fmt.Errorf("error parsing BGG project (%s)", item.Item.Id), unmarshalErr)
			continue
		}

		projects = append(projects, Project{
			Host:  "Board Game Geek",
			Title: bggProject.Title,
			Url:   templ.SafeURL(fmt.Sprintf("https://boardgamegeek.com/boardgame/%s", item.Item.Id)),
			Image: Image{
				Src: bggProject.ImageSrc,
				Alt: fmt.Sprintf("Board Game: %s", bggProject.Title),
			},
			Logo:   components.BGGLogo(),
			Topics: bggProject.Tags,
		})
	}

	return projects, err
}
