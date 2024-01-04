package api

import (
	"fmt"
	"net/http"
	"time"
)

var Now = func() time.Time {
	return time.Now()
}

func GetCopyright(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	year := Now().Year()
	fmt.Fprintf(w, "© %s %d", name, year)
}
