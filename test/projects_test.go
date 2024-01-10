package test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/NDoolan360/go-htmx-site/api"
	"github.com/NDoolan360/go-htmx-site/utils"
)

func TestProjects(t *testing.T) {
	tests := []struct {
		hosts []string
		want  []utils.Project
	}{
		{[]string{"github"}, githubExpected},
		{[]string{"bgg"}, bggExpected},
		{[]string{"cults3d"}, cults3DExpected},
		{[]string{"github", "bgg", "cults3d"}, append(append(githubExpected, bggExpected...), cults3DExpected...)},
	}
	utils.Fetch = func(url string) (string, error) {
		switch url {
		case utils.HostMap["github"].Path:
			return githubMock, nil
		case utils.HostMap["bgg"].Path:
			return bggMock, nil
		case utils.HostMap["cults3d"].Path:
			return cults3DMock, nil
		}
		return "", errors.New("mock url not defined")
	}
	for _, tc := range tests {
		projects, errs := api.FetchProjects(tc.hosts)
		if len(errs) > 0 {
			t.Fatalf("Got error: %v", errs)
		} else if !reflect.DeepEqual(projects, tc.want) {
			t.Fatalf("Got %v;\nwant %v", projects, tc.want)
		}
	}
}
