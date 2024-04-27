package api

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/projects":
		GetProjects(w, r)
	case "/resume":
		GetResume(w, r)
	case "/sitemap.xml":
		GetSitemap(w, r)
	case "/":
		GetIndex(w, r)
	default:
		http.StripPrefix("/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	}
}
