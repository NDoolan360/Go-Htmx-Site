package layouts

import "embed"

//go:embed templates/*.html
//go:embed templates/*.xml
var templates embed.FS
