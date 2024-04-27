package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	fmt.Println("Server listening on port:\thttp://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("http.ListenAndServe(\":3000\", nil) failed with: %v", err)
	}
}
