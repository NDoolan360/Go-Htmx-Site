package main

type Host interface {
	Fetch() ([]byte, error)
	Parse([]byte) ([]Project, error)
}

type GithubHost struct {
	BaseURL string
	User    string
}

type GithubProject struct {
	Title       string   `json:"name"`
	Description string   `json:"description"`
	Url         string   `json:"html_url"`
	Language    string   `json:"language"`
	Topics      []string `json:"topics"`
	Fork        bool     `json:"fork"`
}

type BggHost struct {
	BaseURL  string
	Geeklist string
}

type BggProject struct {
	Item struct {
		Id string `xml:"objectid,attr"`
	} `xml:"item"`
}

type BggItem struct {
	Title    string   `xml:"boardgame>name"`
	ImageSrc string   `xml:"boardgame>image"`
	Tags     []string `xml:"boardgame>boardgamemechanic"`
}

type Cults3dHost struct {
	BaseURL string
	User    string
}

type Cults3dData struct {
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
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}
