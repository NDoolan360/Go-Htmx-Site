package layouts

import "html/template"

// SitemapTemplate represents the data structure for the sitemap.xml template.
type SitemapTemplate struct {
	Url string
}

var Sitemap = template.Must(template.ParseFS(templates, "*/sitemap.xml"))
