package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/NDoolan360/go-htmx-site/website/layouts"
	"github.com/NDoolan360/go-htmx-site/website/pages"
)

func main() {
	for file, Render := range handlers {
		buf := bytes.NewBufferString("")
		Render(buf)
		os.WriteFile(path.Join("static", file), buf.Bytes(), 0644)
	}
}

var handlers = map[string]func(io.Writer){
	"index.html": func(w io.Writer) {
		layouts.BaseLayout(
			"Nathan Doolan",
			"A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
			pages.IndexHeader(),
			pages.IndexMain(fmt.Sprint(time.Now().Year())),
		).Render(context.Background(), w)
	},
	"resume.html": func(w io.Writer) {
		layouts.BaseLayout(
			"Resume",
			"",
			nil,
			pages.Markdown("/content/resume.md"),
		).Render(context.Background(), w)
	},
}
