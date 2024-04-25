package api

import (
	"errors"
	"html/template"
	"reflect"
	"testing"
)

// TestFetchProjects tests the FetchProjects function to ensure it fetches and parses projects correctly.
func TestFetchProjects(t *testing.T) {
	tests := []struct {
		name  string
		hosts []string
		want  []Project
	}{
		{"Github Projects Test", []string{"github"}, GithubExpected},
		{"BGG Projects Test", []string{"bgg"}, BggExpected},
		{"Cults3D Projects Test", []string{"cults3d"}, Cults3DExpected},
		{"All Projects Test", []string{"github", "bgg", "cults3d"}, append(append(GithubExpected, BggExpected...), Cults3DExpected...)},
	}

	// mock fetching contents
	Fetch = func(request *Request) ([]byte, error) {
		switch request.Path {
		case "https://api.github.com/users/NDoolan360/repos?sort=stars":
			return []byte(githubMock), nil
		case "https://boardgamegeek.com/xmlapi/geeklist/332832":
			return []byte(bggMock), nil
		case "https://boardgamegeek.com/xmlapi/boardgame/330653":
			return []byte(bggXMLMock), nil
		case "https://cults3d.com/graphql":
			return []byte(cults3DMock), nil
		default:
			return nil, errors.New("mock url not defined")
		}
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			projects, errs := FetchProjects(tc.hosts)
			if errs != nil {
				t.Fatalf("Got error: %v", errs)
			}
			if len(tc.want) != len(projects) {
				t.Fatalf("Expected projects to be len(%d), Got len(%d)", len(tc.want), len(projects))
			}
			for i, project := range projects {
				if tc.want[i].Host != project.Host {
					t.Fatalf("Host\nReceived: %v;\nExpected: %v", project.Host, tc.want[i].Host)
				}
				if tc.want[i].Title != project.Title {
					t.Fatalf("Title\nReceived: %v;\nExpected: %v", project.Title, tc.want[i].Title)
				}
				if !reflect.DeepEqual(tc.want[i].Description, project.Description) {
					t.Fatalf("Description\nReceived: %v;\nExpected: %v", project.Description, tc.want[i].Description)
				}
				if tc.want[i].Url != project.Url {
					t.Fatalf("Url\nReceived: %v;\nExpected: %v", project.Url, tc.want[i].Url)
				}
				if !reflect.DeepEqual(tc.want[i].Image, project.Image) {
					t.Fatalf("Image\nReceived: %v;\nExpected: %v", project.Image, tc.want[i].Image)
				}
				if !reflect.DeepEqual(tc.want[i].Language, project.Language) {
					t.Fatalf("Image\nReceived: %v;\nExpected: %v", project.Language, tc.want[i].Image)
				}
				if !reflect.DeepEqual(tc.want[i].Logo, project.Logo) {
					t.Fatalf("Logo\nReceived: %v;\nExpected: %v", project.Logo, tc.want[i].Logo)
				}
				if !reflect.DeepEqual(tc.want[i].Topics, project.Topics) {
					t.Fatalf("Topics\nReceived: %v;\nExpected: %v", project.Topics, tc.want[i].Topics)
				}
			}
		})
	}
}

var githubMock = `
[
    {
        "name": "Test",
        "html_url": "https://github.com/NDoolan360/Test",
        "description": "My hand crafted Test",
        "fork": false,
        "language": "Go",
        "topics": [
            "test1",
            "test2",
            "test3"
        ]
    },
    {
        "name": "Forked-Test",
        "html_url": "https://github.com/NDoolan360/Forked-Test",
        "description": null,
        "fork": true,
        "language": "TypeScript",
        "topics": []
    },
    {
        "name": "No-Topics-Test",
        "html_url": "https://github.com/NDoolan360/No-Topics-Test",
        "description": "Just an empty husk without topics",
        "fork": false,
        "language": "Rust",
        "topics": []
    }
]`

var bggMock = `
<geeklist>
    <item objectid="330653" objectname="Cake Toppers"></item>
</geeklist>`

var bggXMLMock = `
<boardgames termsofuse="https://boardgamegeek.com/xmlapi/termsofuse">
    <boardgame objectid="330653">
        <name primary="true" sortindex="1">Cake Toppers</name>
        <image>https://cf.geekdo-images.com/wFwQ-MEGf6BLIyV77hQvHQ__original/img/jGDJHygR3da__4gT0pMzKAD1SQU=/0x0/filters:format(png)/pic5982841.png</image>
        <boardgamemechanic objectid="2040">Hand Management</boardgamemechanic>
        <boardgamemechanic objectid="2914">Increase Value of Unchosen Resources</boardgamemechanic>
        <boardgamemechanic objectid="2048">Pattern Building</boardgamemechanic>
        <boardgamemechanic objectid="2004">Set Collection</boardgamemechanic>
    </boardgame>
</boardgames>`

var cults3DMock = `
{
  "data": {
      "user": {
          "creations": [
              {
                  "name": "Reciprocating Rack and Pinion Fidget V2",
                  "url": "https://cults3d.com/en/3d-model/gadget/reciprocating-rack-and-pinion-fidget-v2",
                  "illustrationImageUrl": "https://files.cults3d.com/uploaders/20027643/illustration-file/5371a13c-5cfa-4ce7-aebb-aedfa3865bd1/RRaPv2.png",
                  "tags": [
                      "rack",
                      "pinion",
                      "fidget",
                      "toy",
                      "reciprocating"
                  ]
              },
              {
                  "name": "Thought Processor",
                  "url": "https://cults3d.com/en/3d-model/art/thought-processor",
                  "illustrationImageUrl": "https://files.cults3d.com/uploaders/20027643/illustration-file/9a4f2164-33a2-49ca-8b3b-2975c9bdf03b/RRaP2-Copy.png",
                  "tags": [
                      "bust",
                      "crt",
                      "computer",
                      "monitor",
                      "display",
                      "screen"
                  ]
              }
          ]
      }
  }
}`

var GithubExpected = []Project{
	{
		Host:        "Github",
		Logo:        GetSVGLogo("github"),
		Title:       "Test",
		Description: "My hand crafted Test",
		Url:         "https://github.com/NDoolan360/Test",
		Language: Language{
			Name:   "Go",
			Colour: template.CSS("#00ADD8"),
		},
		Topics: []string{"test1", "test2", "test3"},
	},
}

var BggExpected = []Project{
	{
		Host:  "Board Game Geek",
		Title: "Cake Toppers",
		Url:   "https://boardgamegeek.com/boardgame/330653",
		Image: Image{
			Src: template.HTMLAttr(`src="https://cf.geekdo-images.com/wFwQ-MEGf6BLIyV77hQvHQ__original/img/jGDJHygR3da__4gT0pMzKAD1SQU=/0x0/filters:format(png)/pic5982841.png"`),
			Alt: template.HTMLAttr(`alt="Board Game: Cake Toppers"`),
		},
		Logo:   GetSVGLogo("bgg"),
		Topics: []string{"Hand Management", "Increase Value of Unchosen Resources", "Pattern Building", "Set Collection"},
	},
}

var Cults3DExpected = []Project{
	{
		Host:  "Cults3D",
		Title: "Reciprocating Rack and Pinion Fidget V2",
		Url:   "https://cults3d.com/en/3d-model/gadget/reciprocating-rack-and-pinion-fidget-v2",
		Image: Image{
			Src: template.HTMLAttr(`src="https://files.cults3d.com/uploaders/20027643/illustration-file/5371a13c-5cfa-4ce7-aebb-aedfa3865bd1/RRaPv2.png"`),
			Alt: template.HTMLAttr(`alt="3D Model: Reciprocating Rack and Pinion Fidget V2"`),
		},
		Logo:   GetSVGLogo("cults3d"),
		Topics: []string{"rack", "pinion", "fidget", "toy", "reciprocating"},
	},
	{
		Host:  "Cults3D",
		Title: "Thought Processor",
		Url:   "https://cults3d.com/en/3d-model/art/thought-processor",
		Logo:  GetSVGLogo("cults3d"),
		Image: Image{
			Src: template.HTMLAttr(`src="https://files.cults3d.com/uploaders/20027643/illustration-file/9a4f2164-33a2-49ca-8b3b-2975c9bdf03b/RRaP2-Copy.png"`),
			Alt: template.HTMLAttr(`alt="3D Model: Thought Processor"`),
		},
		Topics: []string{"bust", "crt", "computer", "monitor", "display", "screen"},
	},
}
