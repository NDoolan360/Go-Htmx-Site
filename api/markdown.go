package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// MarkdownTemplate represents the data structure for the markdown.gohtml template.
type MarkdownTemplate struct {
	Title           string
	Description     string
	MarkdownSource  string
	MarkdownSrcAttr template.HTMLAttr
}

// GetIndex handles the request for rendering the index page.
func GetMarkdown(w http.ResponseWriter, r *http.Request) {
	filesrc := r.URL.Query()["file"]
	if len(filesrc) < 1 {
		http.Error(w, "No file provided.", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles(
		GetApiAsset("template/markdown.gohtml"),
		GetApiAsset("template/head.gohtml"),
		GetApiAsset("template/theme-switch.gohtml"),
	))

	execErr := tmpl.Execute(w, MarkdownTemplate{
		Title:           "",
		Description:     "",
		MarkdownSource:  filesrc[0],
		MarkdownSrcAttr: template.HTMLAttr(fmt.Sprintf(`src="%s"`, filesrc[0])),
	})
	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
	}
}
