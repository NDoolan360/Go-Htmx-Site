package utils

import "html/template"

// Index.gohtml

type Index struct {
	Title          string
	Description    string
	InternalLinks  []Link
	ExternalLinks  []Link
	ThemeSwitchSVG template.HTML
	Profile
	Experiences  []Experience
	ToolSections []ToolSection
	Copyright    string
}

// Projects.gohtml

type Project struct {
	Host           string
	LogoSVG        template.HTML
	ImageAttr      []template.HTMLAttr
	Title          string   `json:"name"`
	Description    string   `json:"description"`
	HtmlUrl        string   `json:"html_url"`
	Topics         []string `json:"topics"`
	Fork           bool     `json:"fork"`
	Language       string   `json:"language"`
	LanguageColour template.CSS
}

// Other

type Link struct {
	Label string
	URL   template.URL
	Logo  template.HTML
}

type Profile struct {
	ImageAttr  []template.HTMLAttr
	Paragraphs []string
}

type Experience struct {
	Date
	Workplace Link
	Positions []Position
	Topics    []string
	Education bool
}

type Date struct {
	Start string
	End   string
}

type Position struct {
	Title   string
	Current bool
}

type ToolSection struct {
	Title string
	Links []Link
}
