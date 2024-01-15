package main

import (
	"fmt"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/api"
)

func main() {
	http.HandleFunc("/api/index", http.HandlerFunc(api.GetIndex))
	http.HandleFunc("/api/projects", http.HandlerFunc(api.GetProjects))
	http.HandleFunc("/api/markdown", http.HandlerFunc(api.GetMarkdown))
	http.HandleFunc("/markdown", http.HandlerFunc(api.GetMarkdown)) 	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			api.GetIndex(w, r)
		} else {
			http.StripPrefix("/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
		}
	})

	port := ":3000"
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(port, nil)
}
