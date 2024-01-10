package main

import (
	"fmt"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/api"
)

var apiHandlers = map[string]http.HandlerFunc{
	"/api/index":       http.HandlerFunc(api.GetIndex),
	"/api/projects":    http.HandlerFunc(api.GetProjects),
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		api.GetIndex(w, r)
	} else {
		http.StripPrefix("/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	}
}

func main() {
	for endpoint, handler := range apiHandlers {
		http.HandleFunc(endpoint, handler)
	}
	http.HandleFunc("/", HandleIndex)

	port := ":3000"
	fmt.Printf("Server listening on port %s\n", port)
	panic(http.ListenAndServe(port, nil))
}
