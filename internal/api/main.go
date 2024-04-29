package api

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NDoolan360/go-htmx-site/internal/handlers"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	http.HandleFunc("/projects", handlers.Projects)
	http.HandleFunc("/resume", handlers.Resume)
	http.HandleFunc("/sitemap.xml", handlers.Sitemap)
	http.HandleFunc("/", handlers.Index)

	if _, ok := os.LookupEnv("NETLIFY"); ok {
		lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
	}
}
