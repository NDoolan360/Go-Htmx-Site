package layouts

import (
	"html/template"

	"github.com/NDoolan360/go-htmx-site/internal/components"
)

// IndexTemplate represents the data structure for the index.html template.
type IndexTemplate struct {
	Title         string
	Description   string
	InternalLinks []components.Link
	ExternalLinks []components.Link
	Profile       components.Profile
	Experiences   []components.Experience
	ToolSections  []components.ToolSection
	Copyright     string
}

var Index = template.Must(template.ParseFS(templates, "*/index.html", "*/head.html"))
