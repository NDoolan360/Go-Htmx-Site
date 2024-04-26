package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// GetSitemap handles the request for rendering the sitemap.
func GetSitemap(w http.ResponseWriter, r *http.Request) {
	markdownTemplate := template.Must(template.ParseFiles("templates/sitemap.xml.tmpl"))

	execErr := markdownTemplate.Execute(w, SitemapTemplate{
		Url: fmt.Sprintf("https://%s", r.Host),
	})

	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
	}
}
