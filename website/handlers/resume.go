package handlers

import (
	"net/http"

	"github.com/NDoolan360/go-htmx-site/website/layouts"
	"github.com/NDoolan360/go-htmx-site/website/pages"
)

// Resume handles the request for rendering the resume page.
func Resume(w http.ResponseWriter, r *http.Request) {
	layouts.BaseLayout("Resume", "", nil, pages.Markdown("Resume.md")).Render(r.Context(), w)
}
