package api

import (
	"html/template"
	"net/http"
)

// GetIndex handles the request for rendering the index page.
func TypesHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, nil)
}

// IndexTemplate represents the data structure for the index.gohtml template.
type IndexTemplate struct {
	// Basic page information for <head>
	Title       string
	Description string

	// Navigation links
	InternalLinks []Link
	ExternalLinks []Link

	// Profile information
	Profile Profile

	// Professional and educational experiences
	Experiences []Experience

	// Tools and technologies used section
	ToolSections []ToolSection

	// Copyright string
	Copyright string
}

// ProjectsTemplate represents the data structure for the projects.gohtml template.
type ProjectsTemplate struct {
	Projects []Project
}

// MarkdownTemplate represents the data structure for the markdown.gohtml template.
type MarkdownTemplate struct {
	Title           string
	Description     string
	MarkdownSource  string
	MarkdownSrcAttr template.HTMLAttr
}

// Link represents a hyperlink with label, URL, and optional logo.
type Link struct {
	Label string
	URL   template.URL
	Logo  template.HTML
}

// Profile represents user profile information with image and paragraphs.
type Profile struct {
	ImageAttr  []template.HTMLAttr
	Paragraphs []string
}

// Experience represents a work experience or education entry.
type Experience struct {
	Date
	Link
	Positions []Position
	Knowledge []Link
	Education bool
}

// Date represents a start and end date for an experience.
type Date struct {
	Start string
	End   string
}

// Position represents a job title and current job status for work experience.
type Position struct {
	Title   string
	Current bool
}

// ToolSection represents a section with a title and links to tools used.
type ToolSection struct {
	Title string
	Links []Link
}

// Project represents information about a personal project.
type Project struct {
	Host  string
	Title string
	Description template.HTML
	Url   string
	Image
	Language
	Logo        template.HTML
	Topics      []string
}

type Image struct {
	Src template.HTMLAttr
	Alt template.HTMLAttr
}

type Language struct {
	Name   string
	Colour template.CSS
}

type Host struct {
	Request
	Parser func([]byte) ([]Project, error)
}

type Request struct {
	Method      string
	Path        string
	Username    string
	Password    string
	ContentType string
	Body        string
}
