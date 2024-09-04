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

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var wg sync.WaitGroup
	buf := bytes.NewBufferString("")

	for _, hostKey := range request.MultiValueQueryStringParameters["host"] {
		host, ok := hostMap[hostKey]
		if !ok {
			return nil, fmt.Errorf("Interface for host '%s' not found.", host)
		}

		wg.Add(1)
		go func(hostKey string, host Host, ctx context.Context, buf io.Writer, wg *sync.WaitGroup) {
			defer wg.Done()

			data, err := host.Fetch()
			if err != nil {
				log.Print(hostKey, ": ", err)
				return
			}

			projects, err := host.Parse(data)
			if err != nil {
				log.Print(hostKey, ": ", err)
				return
			}

			for _, project := range projects {
				err := ProjectTemplate(project).Render(ctx, buf)
				if err != nil {
					log.Print(hostKey, ": ", err)
				}
			}
		}(hostKey, host, ctx, buf, &wg)
	}

	wg.Wait()

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "text/html; charset=utf-8"},
		Body:       buf.String(),
	}, nil
}
