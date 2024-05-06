module github.com/NDoolan360/go-htmx-site/api/experience

go 1.22.2

require (
	github.com/a-h/templ v0.2.663
	github.com/aws/aws-lambda-go v1.47.0
)

require github.com/NDoolan360/go-htmx-site/website v0.0.0-20240506034505-d3c61fad215a // indirect

replace github.com/NDoolan360/go-htmx-site/website => ../../website
