package components

import "embed"

//go:embed templates/*.html
var templates embed.FS

//go:embed logos/*.svg
var logos embed.FS
