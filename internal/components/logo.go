package components

import (
	"html/template"
)

// GetSVGLogo reads an SVG logo file and returns it as an HTML template.
func GetSVGLogo(logo string) template.HTML {
	svg, err := logos.ReadFile("logos/" + logo + ".svg")
	if err != nil {
		panic(err)
	}
	return template.HTML(string(svg))
}
