package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"log"

	"github.com/a-h/templ"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//go:generate make templates -C ../..

type Experience struct {
	DateStart string        `json:"dateStart"`
	DateEnd   string        `json:"dateEnd"`
	Location  string        `json:"location"`
	Positions []Position    `json:"positions"`
	Link      templ.SafeURL `json:"link"`
	Topics    []Topic       `json:"topics"`
}

type Position struct {
	Role   string `json:"role"`
	Active bool   `json:"active"`
}

type Topic struct {
	Label string        `json:"label"`
	Link  templ.SafeURL `json:"link"`
}

//go:embed experience.json
var experienceFile []byte

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	buf := bytes.NewBufferString("")

	var experiences []Experience
	err := json.Unmarshal(experienceFile, &experiences)
	if err != nil {
		log.Print(err)
	}

	for _, experience := range experiences {
		err := ExperienceTemplate(experience).Render(ctx, buf)
		if err != nil {
			log.Print(err)
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/html; charset=utf-8"},
		Body:       buf.String(),
	}, nil
}
