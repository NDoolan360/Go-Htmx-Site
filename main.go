package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/api"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/projects":
		api.GetProjects(w, r)
	case "/resume":
		api.GetResume(w, r)
	case "/sitemap.xml":
		api.GetSitemap(w, r)
	case "/":
		api.GetIndex(w, r)
	default:
		http.StripPrefix("/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port:\thttp://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("http.ListenAndServe(\":3000\", nil) failed with: %v", err)
	}
}
