package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// GetIndex handles the request for rendering the index page.
func GetMarkdown(w http.ResponseWriter, r *http.Request) {
	fileSource := r.URL.Query()["file"]
	if len(fileSource) < 1 {
		http.Error(w, "No file provided.", http.StatusBadRequest)
		return
	}

	markdownTemplate := template.Must(template.ParseFiles(
		GetApiAsset("template/markdown.html.tmpl"),
		GetApiAsset("template/head.html.tmpl"),
		GetApiAsset("template/theme-switch.html.tmpl"),
	))

	execErr := markdownTemplate.Execute(w, MarkdownTemplate{
		Title:           fileSource[0],
		Description:     "",
		MarkdownSource:  fileSource[0],
		MarkdownSrcAttr: template.HTMLAttr(fmt.Sprintf(`src="%s"`, fileSource[0])),
	})
	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
	}
}
