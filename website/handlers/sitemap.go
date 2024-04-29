package handlers

import (
	"fmt"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/website/pages"
)

// Sitemap handles the request for rendering the sitemap.
func Sitemap(w http.ResponseWriter, r *http.Request) {
	pages.Sitemap(fmt.Sprintf("https://%s", r.Host)).Render(r.Context(), w)
}
