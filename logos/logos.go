package logos

import (
	"embed"
	"html/template"
)

//go:embed assets/*.svg
var assets embed.FS

// GetSVGLogo reads an SVG logo file and returns it as an HTML template.
func GetSVGLogo(logo string) template.HTML {
	svg, err := assets.ReadFile("assets/" + logo + ".svg")
	if err != nil {
		panic(err)
	}
	return template.HTML(string(svg))
}
