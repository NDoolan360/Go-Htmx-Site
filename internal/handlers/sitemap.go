package handlers

import (
	"fmt"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/internal/layouts"
)

// Sitemap handles the request for rendering the sitemap.
func Sitemap(w http.ResponseWriter, r *http.Request) {
	layouts.Sitemap.Execute(w, layouts.SitemapTemplate{
		Url: fmt.Sprintf("https://%s", r.Host),
	})
}
