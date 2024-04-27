package components

import (
	"html/template"
)

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
	Link
	Positions []Position
	Knowledge []Link
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

type Image struct {
	Src template.HTMLAttr
	Alt template.HTMLAttr
}

type Language struct {
	Name   string
	Colour template.CSS
}
