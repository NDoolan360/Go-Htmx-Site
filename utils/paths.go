package utils

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func FromPWD(path string) string {

	pwd, _ := os.Getwd()
	currentDirBase := filepath.Base(pwd)
	switch currentDirBase {
	case "test":
		return "../" + path
	case "api":
		if _, inVercel := os.LookupEnv("VERCEL"); inVercel {
			return strings.TrimPrefix(path, "api/")
		} else {
			return "../" + path
		}
	default:
		return path
	}
}

func GetSVGLogo(logo string) template.HTML {
	svg, err := os.ReadFile(FromPWD("api/logos/" + logo + ".svg"))
	if err != nil {
		panic(err)
	}
	return template.HTML(svg)
}
