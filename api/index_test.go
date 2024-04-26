package api

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// TestCopyright tests the Copyright function to ensure it generates the correct copyright string.
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
		Now = func() time.Time { return tc.now }
		out := Copyright(tc.name)
		if diff := cmp.Diff(tc.want, out); diff != "" {
			t.Errorf("unexpected copyright (-want +got):\n%s", diff)
		}
	}
}
