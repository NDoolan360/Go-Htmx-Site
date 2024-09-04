package main

type Host interface {
	Fetch() ([]byte, error)
	Parse([]byte) ([]Project, error)
}

var hostMap = map[string]Host{
	"github":  GithubHost{BaseURL: "https://api.github.com", User: "NDoolan360"},
	"bgg":     BggHost{BaseURL: "https://boardgamegeek.com/xmlapi", Geeklist: "332832"},
	"cults3d": Cults3dHost{BaseURL: "https://cults3d.com", User: "ND360"},
}
