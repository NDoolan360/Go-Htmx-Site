package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Method   string
	Path     string
	Username string
	Password string
}

type Outcome struct {
	Host   string `json:"host"`
	Status string `json:"status"`
}

var dependencyHealthEndpoints = map[string]Request{
	"github": {
		Method: "GET",
		Path:   "https://api.github.com",
	},
	"cults3d": {
		Method:   "POST",
		Path:     "https://cults3d.com/graphql?query=%7B__typename%7D",
		Username: os.Getenv("CULTS3D_USERNAME"),
		Password: os.Getenv("CULTS3D_API_KEY"),
	},
	"bgg": {
		Method: "GET",
		Path:   "https://boardgamegeek.com/xmlapi/search",
	},
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var wg sync.WaitGroup
	dependencies := []Outcome{}

	for host, req := range dependencyHealthEndpoints {
		wg.Add(1)

		go func(host string, req Request, wg *sync.WaitGroup) {
			defer wg.Done()
			status, err := fetchStatus(req)

			var outcome Outcome
			if err != nil {
				log.Print(host, ": ", err)
				outcome = Outcome{Host: host, Status: err.Error()}
			} else {
				outcome = Outcome{Host: host, Status: status}
			}
			dependencies = append(dependencies, outcome)
		}(host, req, &wg)
	}

	wg.Wait()

	status := map[string]interface{}{"host": "n.doolan.dev", "status": "200 OK", "dependencies": dependencies}
	body, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(body),
	}, nil
}

func fetchStatus(request Request) (string, error) {
	outgoingRequest, err := http.NewRequest(request.Method, request.Path, nil)
	if err != nil {
		return "", err
	}
	if len(request.Username) > 0 && len(request.Password) > 0 {
		outgoingRequest.SetBasicAuth(request.Username, request.Password)
	}

	client := &http.Client{}
	response, err := client.Do(outgoingRequest)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	return response.Status, nil
}
