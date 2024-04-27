package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NDoolan360/go-htmx-site/internal/layouts"
	"github.com/NDoolan360/go-htmx-site/internal/pages"
)

// Index handles the request for rendering the index page.
func Index(w http.ResponseWriter, r *http.Request) {
	year := time.Now().Year()

	layouts.BaseLayout(
		"Nathan Doolan",
		"A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		pages.IndexHeader(),
		pages.IndexMain(fmt.Sprint(year)),
	).Render(r.Context(), w)
}
