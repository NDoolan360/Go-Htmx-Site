package logos

import (
	"html/template"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestGetSVGLogo tests the GetSVGLogo function to ensure it fetches logos correctly.
func TestGetSVGLogo(t *testing.T) {
	want := template.HTML("<svg viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" id=\"logo\" class=\"logo\"\r\n    aria-label=\"Board Game Geek Logo\">\r\n    <path fill=\"#ff5100\"\r\n        d=\"M 16.870973,5.6604271 19.008331,1 5.1505932,6.1006591 5.9080939,12.209966 4.628579,13.443659 8.443343,23 l 8.109201,-2.987661 2.818877,-6.617401 -1.210493,-1.166992 0.892644,-7.1620351 z\"\r\n    />\r\n</svg>")
	out := GetSVGLogo("bgg")
	if diff := cmp.Diff(want, out); diff != "" {
		t.Errorf("unexpected GetSVGLogo (-want +got):\n%s", diff)
	}
}
