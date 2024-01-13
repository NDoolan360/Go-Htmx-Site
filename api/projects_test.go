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
		hosts []string
		want  []Project
	}{
		{[]string{"github"}, GithubExpected},
		{[]string{"bgg"}, BggExpected},
		{[]string{"cults3d"}, Cults3DExpected},
		{[]string{"github", "bgg", "cults3d"}, append(append(GithubExpected, BggExpected...), Cults3DExpected...)},
	}
	Fetch = mockFetch
	for _, tc := range tests {
		projects, errs := FetchProjects(tc.hosts)
		if len(errs) > 0 {
			t.Fatalf("Got error: %v", errs)
		} else if !reflect.DeepEqual(projects, tc.want) {
			t.Fatalf("Got %v;\nwant %v", projects, tc.want)
		}
	}
}

var GithubExpected = []Project{
	{
		Host:           "Github",
		LogoSVG:        GetSVGLogo("github"),
		Title:          "Test",
		Description:    "My hand crafted Test",
		HtmlUrl:        "https://github.com/NDoolan360/Test",
		Topics:         []string{"test1", "test2", "test3"},
		Fork:           false,
		Language:       "Go",
		LanguageColour: "#00ADD8",
	},
}

var BggExpected = []Project{
	{
		Host:    "Board Game Geek",
		LogoSVG: GetSVGLogo("bgg"),
		ImageAttr: []template.HTMLAttr{
			template.HTMLAttr(`src="https://cf.geekdo-images.com/wFwQ-MEGf6BLIyV77hQvHQ__micro/img/qOEv3ACF09F-_zGh0cSMIOXQrVs=/fit-in/64x64/filters:strip_icc()/pic5982841.png"`),
			template.HTMLAttr(`alt="Board Game: Cake Toppers"`),
		},
		Title:       "Cake Toppers",
		Description: "Bakers assemble the most outrageous cakes to top each other.",
		HtmlUrl:     "https://boardgamegeek.com/boardgame/330653/cake-toppers",
	},
}

var Cults3DExpected = []Project{
	{
		Host:    "Cults3D",
		LogoSVG: GetSVGLogo("cults3d"),
		ImageAttr: []template.HTMLAttr{
			template.HTMLAttr(`src="https://files.cults3d.com/uploaders/20027643/illustration-file/5371a13c-5cfa-4ce7-aebb-aedfa3865bd1/RRaPv2.png"`),
			template.HTMLAttr(`alt="RRaPv2.png Reciprocating Rack and Pinion Fidget V2"`),
		},
		Title:   "Reciprocating Rack and Pinion Fidget V2",
		HtmlUrl: "https://cults3d.com/en/3d-model/gadget/reciprocating-rack-and-pinion-fidget-v2",
	},
	{
		Host:    "Cults3D",
		LogoSVG: GetSVGLogo("cults3d"),
		ImageAttr: []template.HTMLAttr{
			template.HTMLAttr(`src="https://files.cults3d.com/uploaders/20027643/illustration-file/9a4f2164-33a2-49ca-8b3b-2975c9bdf03b/RRaP2-Copy.png"`),
			template.HTMLAttr(`alt="RRaP2-Copy.png Thought Processor"`),
		},
		Title:   "Thought Processor",
		HtmlUrl: "https://cults3d.com/en/3d-model/art/thought-processor",
	},
}

var mockFetch = func(url string) (string, error) {
	switch url {
	case "https://api.github.com/users/NDoolan360/repos?sort=stars":
		return githubMock, nil
	case "https://boardgamegeek.com/geeksearch.php?action=search&advsearch=1&objecttype=boardgame&include%5Bdesignerid%5D=133893":
		return bggMock, nil
	case "https://cults3d.com/en/users/ND360/3d-models":
		return cults3DMock, nil
	}
	return "", errors.New("mock url not defined")
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
<table>
    <tbody>
        <tr>
            <th class="collection_thumbnail">
                <span class="sr-only">Thumbnail image</span>
            </th>
            <th>
                <a>Title</a>
            </th>
        </tr>
        <tr id="row_">
            <td class="collection_thumbnail">
                <a href="/boardgame/330653/cake-toppers">
                    <img alt="Board Game: Cake Toppers" src="https://cf.geekdo-images.com/wFwQ-MEGf6BLIyV77hQvHQ__micro/img/qOEv3ACF09F-_zGh0cSMIOXQrVs=/fit-in/64x64/filters:strip_icc()/pic5982841.png" />
                </a>
            </td>
            <td>
                <div>
                    <a class="primary">Cake Toppers</a>
                </div>
                <p class="smallefont">Bakers assemble the most outrageous cakes to top each other.</p>
            </td>
        </tr>
    </tbody>
</table>`

var cults3DMock = `
<article class="crea">
  <div>
    <a title="Reciprocating Rack and Pinion Fidget V2"
      href="/en/3d-model/gadget/reciprocating-rack-and-pinion-fidget-v2">
      <div>
        <picture>
          <img alt="RRaPv2.png Reciprocating Rack and Pinion Fidget V2"
            data-src="https://images.cults3d.com/PFIDNlM1rYYHDszVySD-6bg0sJk=/246x246/filters:no_upscale()/https://files.cults3d.com/uploaders/20027643/illustration-file/5371a13c-5cfa-4ce7-aebb-aedfa3865bd1/RRaPv2.png">
        </picture>
      </div>
      <div>
        <h3>Reciprocating Rack and Pinion Fidget V2</h3>
      </div>
    </a>
  </div>
</article>
<article class="crea">
  <div>
    <a title="Thought Processor" href="/en/3d-model/art/thought-processor">
      <div>
        <picture>
          <img alt="RRaP2-Copy.png Thought Processor"
            data-src="https://images.cults3d.com/BwnOBlJICQURW_aO68cA2AzELzA=/246x246/filters:no_upscale()/https://files.cults3d.com/uploaders/20027643/illustration-file/9a4f2164-33a2-49ca-8b3b-2975c9bdf03b/RRaP2-Copy.png">
        </picture>
      </div>
      <div>
        <h3>Thought Processor</h3>
      </div>
    </a>
  </div>
</article>`
