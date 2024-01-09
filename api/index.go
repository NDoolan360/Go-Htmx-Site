package api

import (
	"fmt"
	"html/template"
	"net/http"

	utils "github.com/NDoolan360/go-htmx-site/src"
)

type Index struct {
	Title         string
	Description   string
	CopyrightYear string
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(
		utils.GetTemplatePath("index.gohtml"),
	)).Execute(w, Index{
		Title:         "Nathan Doolan",
		Description:   "A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		CopyrightYear: fmt.Sprint(Now().Year()),
	})
}
