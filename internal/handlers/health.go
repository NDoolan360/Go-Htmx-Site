package handlers

import (
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, "{\n\tstatus:\"OK\"\n}")
}
