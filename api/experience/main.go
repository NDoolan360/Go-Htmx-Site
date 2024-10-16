package main

import (
	"bytes"
	"context"
	"log"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	buf := bytes.NewBufferString("")

	var experiences []Experience

	workLimit, workOk := request.QueryStringParameters["work"]
	if workOk {
		workLimit, err := strconv.Atoi(workLimit)
		if err != nil || workLimit < 0 {
			return &events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}
		workLimit = min(workLimit, len(workExperiences))
		experiences = append(experiences, workExperiences[:workLimit]...)
	}

	educationParam, eduOk := request.QueryStringParameters["education"]
	if eduOk {
		educationLimit, err := strconv.Atoi(educationParam)
		if err != nil || educationLimit < 0 {
			return &events.APIGatewayProxyResponse{StatusCode: 400}, nil
		}
		educationLimit = min(educationLimit, len(educationExperiences))
		experiences = append(experiences, educationExperiences[:educationLimit]...)
	}

	if !workOk && !eduOk {
		return &events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	for _, experience := range experiences {
		err := ExperienceTemplate(experience).Render(ctx, buf)
		if err != nil {
			log.Print(experience.Location.Name, ": ", err)
			return &events.APIGatewayProxyResponse{StatusCode: 500}, nil
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/html; charset=utf-8"},
		Body:       buf.String(),
	}, nil
}
