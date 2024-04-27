package layouts

import (
	"html/template"
)

// MarkdownTemplate represents the data structure for the markdown.html template.
type MarkdownTemplate struct {
	Title           string
	Description     string
	MarkdownSource  string
	MarkdownSrcAttr template.HTMLAttr
}

var Markdown = template.Must(template.ParseFS(templates, "*/markdown.html", "*/head.html"))
