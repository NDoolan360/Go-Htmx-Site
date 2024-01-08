package api

import (
	"fmt"
	"html/template"
	"net/http"

	utils "github.com/NDoolan360/go-htmx-site/src"
)

type Base struct {
	Title       string
	Description string
}

type Index struct {
	Base
	CopyrightYear string
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		utils.GetTemplate("base.gohtml"),
		utils.GetTemplate("header.gohtml"),
		utils.GetTemplate("index.gohtml"),
	))
	tmpl.Execute(w, Index{
		Base{
			"Nathan Doolan",
			"A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		},
		fmt.Sprint(Now().Year()),
	})
}
