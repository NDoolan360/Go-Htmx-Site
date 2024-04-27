package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// GetIndex handles the request for rendering the index page.
func GetResume(w http.ResponseWriter, r *http.Request) {
	markdownTemplate := template.Must(template.ParseFiles(
		"templates/markdown.html.tmpl",
		"templates/head.html.tmpl",
		"templates/theme-switch.html.tmpl",
	))

	execErr := markdownTemplate.Execute(w, MarkdownTemplate{
		Title:           "Resume",
		Description:     "",
		MarkdownSource:  "Resume.md",
		MarkdownSrcAttr: template.HTMLAttr(fmt.Sprintf(`src="%s"`, "Resume.md")),
	})

	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
	}
}
