package api

import (
	"fmt"
	"html/template"
	"net/http"
)

// GetManifest handles the request for rendering the manifest.
func GetManifest(w http.ResponseWriter, r *http.Request) {
	markdownTemplate := template.Must(template.ParseFiles(
		GetApiAsset("template/manifest.json.tmpl"),
	))

	execErr := markdownTemplate.Execute(w, ManifestTemplate{
		StartUrl: fmt.Sprintf("https://%s/", r.Host),
	})
	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
	}
}
