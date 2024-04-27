package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/internal/layouts"
)

// Resume handles the request for rendering the resume page.
func Resume(w http.ResponseWriter, r *http.Request) {
	layouts.Markdown.Execute(w, layouts.MarkdownTemplate{
		Title:           "Resume",
		Description:     "",
		MarkdownSource:  "Resume.md",
		MarkdownSrcAttr: template.HTMLAttr(fmt.Sprintf(`src="%s"`, "Resume.md")),
	})
}
