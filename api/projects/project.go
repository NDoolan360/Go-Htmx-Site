package main

import "github.com/a-h/templ"

type Project struct {
	Host        string
	Title       string
	Description string
	Url         templ.SafeURL
	Image       Image
	Language    Language
	Logo        templ.Component
	Topics      []string
}

type Image struct {
	Src string
	Alt string
}

type Language struct {
	Name   string
	Colour string
}
