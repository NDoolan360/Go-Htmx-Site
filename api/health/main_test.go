package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

// TestHandler returns HTTP status code 200 in the response body for a health check
func TestHandler(t *testing.T) {
	ctx := context.Background()
	response, handlerErr := handler(ctx, events.APIGatewayProxyRequest{})
	if handlerErr != nil {
		t.Errorf("expected no error but got %s", handlerErr.Error())
	}

	if response.StatusCode != 200 {
		t.Errorf("expected status code 200 but got %d", response.StatusCode)
	}

	var res map[string]interface{}
	unmarshalErr := json.Unmarshal([]byte(response.Body), &res)
	if unmarshalErr != nil {
		t.Errorf("expected no error but got %s", unmarshalErr.Error())
	}

	statusGot := res["status"].(string)
	if statusGot != "200 OK" {
		t.Errorf("expected output '200 OK' but got '%s'", statusGot)
	}
}
