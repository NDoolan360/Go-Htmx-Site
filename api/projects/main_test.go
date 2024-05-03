package main

import (
	"bytes"
	"context"
	_ "embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"
)

//go:embed testdata/expected_github.html
var ExpectedGithubFetchResponse []byte

//go:embed testdata/mock_github.json
var MockGithubHttpResponse []byte

//go:embed testdata/expected_bgg.html
var ExpectedBggFetchResponse []byte

//go:embed testdata/mock_bgg_geeklist.xml
var MockBggHttpResponse []byte

//go:embed testdata/mock_bgg_boardgame.xml
var MockBggXmlHttpResponse []byte

//go:embed testdata/expected_cults3d.html
var ExpectedCults3dFetchResponse []byte

//go:embed testdata/mock_cults3d.json
var MockCults3dHttpResponse []byte

func TestFetchAndParse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/users/NDoolan360/repos":
			w.Write(MockGithubHttpResponse)
		case "/geeklist/332832":
			w.Write(MockBggHttpResponse)
		case "/boardgame/330653":
			w.Write(MockBggXmlHttpResponse)
		case "/graphql":
			w.Write(MockCults3dHttpResponse)
		default:
			t.Errorf("mock url not defined")
		}
	}))

	hostMap = map[string]Host{
		"github":  GithubHost{BaseURL: server.URL, User: "NDoolan360"},
		"bgg":     BggHost{BaseURL: server.URL, Geeklist: "332832"},
		"cults3d": Cults3dHost{BaseURL: server.URL, User: "TEST"},
	}

	tests := []struct {
		name  string
		hosts []string
		want  []byte
	}{
		{"Github Host Test", []string{"github"}, ExpectedGithubFetchResponse},
		{"BGG Host Test", []string{"bgg"}, ExpectedBggFetchResponse},
		{"Cults3D Host Test", []string{"cults3d"}, ExpectedCults3dFetchResponse},
		{"All Projects Test", []string{"github", "bgg", "cults3d"}, bytes.Join([][]byte{ExpectedGithubFetchResponse, ExpectedBggFetchResponse, ExpectedCults3dFetchResponse}, []byte{})},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := handler(context.Background(), events.APIGatewayProxyRequest{MultiValueQueryStringParameters: map[string][]string{"host": tc.hosts}})
			if err != nil {
				t.Fatalf("Got error: %v", err)
			}

			if diff := cmp.Diff(string(tc.want), resp.Body); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}

	server.Close()
}
