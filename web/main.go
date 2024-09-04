package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a-h/templ"

	"github.com/NDoolan360/go-htmx-site/web/templates"
)

var handlers = map[string]templ.Component{
	"index.html": templates.BaseLayout(
		"Nathan Doolan",
		"A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		templates.IndexHeader(),
		templates.IndexMain(fmt.Sprint(time.Now().Year())),
	),
	"resume.html": templates.BaseLayout(
		"Resume",
		"",
		nil,
		templates.Markdown(templ.URL("/content/resume.md")),
	),
}

func main() {
	for file, templComponent := range handlers {
		buf := bytes.NewBufferString("")
		renderErr := templComponent.Render(context.Background(), buf)
		if renderErr != nil {
			log.Fatal(renderErr)
			return
		}

		writeErr := os.WriteFile(file, buf.Bytes(), 0644)
		if writeErr != nil {
			log.Fatal(writeErr)
		} else {
			log.Printf("Generated %s", file)
		}
	}
}
