package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/NDoolan360/go-htmx-site/api/projects/templates"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	hosts := []string{}
	projects, _ := FetchProjects(hosts)

	if len(projects) > 0 {
		for _, project := range projects {
			templates.ProjectTemplate(project)
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            "Hello, world",
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
