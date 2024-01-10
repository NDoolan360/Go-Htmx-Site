package utils

import (
	"html/template"
	"os"
	"path/filepath"
)

func FromPWD(path string) string {
	pwd, _ := os.Getwd()
	currentDirBase := filepath.Base(pwd)
	switch currentDirBase {
	case "test", "api":
		return "../" + path
	}
	return path
}

func GetSVGLogo(logo string) template.HTML {
	svg, err := os.ReadFile(FromPWD("api/logos/" + logo + ".svg"))
	if err != nil {
		panic(err)
	}
	return template.HTML(svg)
}
