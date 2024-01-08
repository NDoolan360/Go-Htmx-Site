package main

import (
	"fmt"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/api"
)

var apiHandlers = map[string]func(http.ResponseWriter, *http.Request){
	"/api/copyright": api.GetCopyright,
	"/api/projects":  api.GetProjects,
}

func main() {
	for endpoint, handler := range apiHandlers {
		http.HandleFunc(endpoint, handler)
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))

	port := ":3000"
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(port, nil)
}
