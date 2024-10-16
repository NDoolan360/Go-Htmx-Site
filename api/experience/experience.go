package main

import "github.com/a-h/templ"

type Experience struct {
	DateStart string
	DateEnd   string
	Location  Place
	Positions []Position
	Link      templ.SafeURL
	Topics    []Topic
}

type Place struct {
	Name string
	Logo templ.Component
}

type Position struct {
	Role   string
	Active bool
}

type Topic struct {
	Label string
	Link  templ.SafeURL
}
