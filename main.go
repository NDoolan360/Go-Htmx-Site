package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/internal/handlers"
)

//go:embed all:static
var static embed.FS

func main() {
	http.HandleFunc("/projects", handlers.Projects)
	http.HandleFunc("/resume", handlers.Resume)
	http.HandleFunc("/sitemap.xml", handlers.Sitemap)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handlers.Index(w, r)
		} else {
			http.StripPrefix("/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
		}
	})

	fmt.Println("Server listening on port:\thttp://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("http.ListenAndServe(\":3000\", nil) failed with: %v", err)
	}
}
