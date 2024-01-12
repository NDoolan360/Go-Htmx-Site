package utils

import (
	"fmt"
	"html/template"
	"os"
	"testing"
	"time"
)

func Copyright(name string) string {
	year := Now().Year()
	return fmt.Sprintf("© %s %d", name, year)
}

// Functions to manipulate runtime behaviour for tests

var Fetch = func(path string) (string, error) {
	return FetchURL(path)
}

var Now = func() time.Time {
	return time.Now()
}

func GetApiResource(path string) string {
	if _, inVercel := os.LookupEnv("VERCEL"); inVercel {
		return path
	} else if testing.Testing() {
		return "../api/" + path
	} else {
		return "api/" + path
	}
}

func GetSVGLogo(logo string) template.HTML {
	svg, err := os.ReadFile(GetApiResource("logos/" + logo + ".svg"))
	if err != nil {
		panic(err)
	}
	return template.HTML(svg)
}
