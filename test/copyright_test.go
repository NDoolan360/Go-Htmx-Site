package test

import (
	"testing"
	"time"

	"github.com/NDoolan360/go-htmx-site/api"
)

func TestCopyright(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want string
	}{
		{
			"Nathan Doolan",
			time.Date(1999, 1, 1, 1, 0, 0, 0, time.UTC),
			"© Nathan Doolan 1999",
		},
		{
			"Nathan Doolan",
			time.Date(1999, 1, 1, 1, 0, 0, 0, time.UTC),
			"© Nathan Doolan 1999",
		},
		{
			"Future Nathan",
			time.Date(2099, 1, 1, 1, 0, 0, 0, time.UTC),
			"© Future Nathan 2099",
		},
	}
	for _, tc := range tests {
		api.Now = func() time.Time { return tc.now }
		out := api.Copyright(tc.name)
		if out != tc.want {
			t.Fatalf("Got %v;\nwant %v", out, tc.want)
		}
	}
}
