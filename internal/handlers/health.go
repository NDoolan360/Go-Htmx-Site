package handlers

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	message, _ := json.Marshal(struct{ status string }{status: "OK"})

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(message)
}
