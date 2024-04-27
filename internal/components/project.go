package components

import (
	"html/template"
)

type ProjectTemplate struct {
	Host        string
	Title       string
	Description template.HTML
	Url         string
	Image
	Language
	Logo   template.HTML
	Topics []string
}

var Project = template.Must(template.ParseFS(templates, "*/project.html"))
