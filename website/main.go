package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/NDoolan360/go-htmx-site/website/layouts"
	"github.com/NDoolan360/go-htmx-site/website/pages"
)

func main() {
	for file, Render := range handlers {
		buf := bytes.NewBufferString("")
		renderErr := Render(buf)
		if renderErr != nil {
			log.Fatal(renderErr)
			return
		}
		writeErr := os.WriteFile(path.Join("static", file), buf.Bytes(), 0644)
		if writeErr != nil {
			log.Fatal(writeErr)
			return
		}
	}
}

var handlers = map[string]func(io.Writer) error{
	"index.html": func(w io.Writer) error {
		return layouts.BaseLayout(
			"Nathan Doolan",
			"A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
			pages.IndexHeader(),
			pages.IndexMain(fmt.Sprint(time.Now().Year())),
		).Render(context.Background(), w)
	},
	"resume.html": func(w io.Writer) error {
		return layouts.BaseLayout(
			"Resume",
			"",
			nil,
			pages.Markdown("/content/resume.md"),
		).Render(context.Background(), w)
	},
}
