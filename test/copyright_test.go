package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/ndoolan360/ndsite/api"
)

func TestGetCopyright(t *testing.T) {
	tests := []struct {
		request http.Request
		now     time.Time
		want    string
	}{
		{
			http.Request{URL: &url.URL{RawQuery: "name=Nathan%20Doolan"}},
			time.Date(1999, 1, 1, 1, 0, 0, 0, time.UTC),
			"© Nathan Doolan 1999",
		},
		{
			http.Request{URL: &url.URL{RawQuery: "name=Nathan Doolan"}},
			time.Date(1999, 1, 1, 1, 0, 0, 0, time.UTC),
			"© Nathan Doolan 1999",
		},
		{
			http.Request{URL: &url.URL{RawQuery: "name=Future%20Nathan"}},
			time.Date(2099, 1, 1, 1, 0, 0, 0, time.UTC),
			"© Future Nathan 2099",
		},
	}
	for _, tc := range tests {
		writer := httptest.NewRecorder()
		api.Now = func() time.Time { return tc.now }
		api.GetCopyright(writer, &tc.request)
		out := writer.Body.String()
		if out != tc.want {
			t.Fatalf("Got %v;\nwant %v", out, tc.want)
		}
	}
}
