module github.com/NDoolan360/go-htmx-site/api/experience

go 1.22.5

require (
	github.com/a-h/templ v0.2.771
	github.com/aws/aws-lambda-go v1.47.0
)

require github.com/NDoolan360/go-htmx-site/web/templates v0.0.0-00010101000000-000000000000

replace github.com/NDoolan360/go-htmx-site/web/templates => ../../web/templates
