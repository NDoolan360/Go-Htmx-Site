package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Host interface {
	Fetch() ([]byte, error)
	Parse([]byte) ([]Project, error)
}

func main() {
	lambda.Start(handler)
}

var hostMap = map[string]Host{
	"github":  GithubHost{BaseURL: "https://api.github.com", User: "NDoolan360"},
	"bgg":     BggHost{BaseURL: "https://boardgamegeek.com/xmlapi", Geeklist: "332832"},
	"cults3d": Cults3dHost{BaseURL: "https://cults3d.com", User: "ND360"},
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var wg sync.WaitGroup
	buf := bytes.NewBufferString("")

	for _, host := range request.MultiValueQueryStringParameters["host"] {
		host, ok := hostMap[host]
		if !ok {
			return nil, fmt.Errorf("Interface for host '%s' not found.", host)
		}

		wg.Add(1)
		go func(host Host, ctx context.Context, buf io.Writer, wg *sync.WaitGroup) {
			defer wg.Done()

			data, err := host.Fetch()
			if err != nil {
				return
			}

			projects, err := host.Parse(data)
			if err != nil {
				return
			}

			for _, project := range projects {
				err := ProjectTemplate(project).Render(ctx, buf)
				if err != nil {
					log.Print(err)
				}
			}
		}(host, ctx, buf, &wg)
	}

	wg.Wait()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/html; charset=utf-8"},
		Body:       buf.String(),
	}, nil
}
