module github.com/NDoolan360/go-htmx-site/api/projects

go 1.22.5

require (
	github.com/NDoolan360/github-langs-go v1.0.0
	github.com/NDoolan360/go-htmx-site/web/templates v0.0.0-00010101000000-000000000000
	github.com/a-h/templ v0.2.771
	github.com/aws/aws-lambda-go v1.47.0
	github.com/google/go-cmp v0.6.0
)

require github.com/pelletier/go-toml/v2 v2.1.1 // indirect

replace github.com/NDoolan360/go-htmx-site/web/templates => ../../web/templates
