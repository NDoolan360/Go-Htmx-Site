package test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/NDoolan360/go-htmx-site/api"
)

func TestProjects(t *testing.T) {
	tests := []struct {
		host []string
		want []api.Project
	}{
		{[]string{"github"}, githubExpected},
		{[]string{"bgg"}, bggExpected},
		{[]string{"cults3d"}, cults3DExpected},
		{[]string{"github", "bgg", "cults3d"}, append(append(githubExpected, bggExpected...), cults3DExpected...)},
	}
	api.Fetch = func(url string) (string, error) {
		switch url {
		case "https://api.github.com/users/NDoolan360/repos?sort=stars":
			return githubMock, nil
		case "https://boardgamegeek.com/geeksearch.php?action=search&advsearch=1&objecttype=boardgame&include%5Bdesignerid%5D=133893":
			return bggMock, nil
		case "https://cults3d.com/en/users/ND360/3d-models":
			return cults3DMock, nil
		}
		return "", errors.New("mock url not defined")
	}
	for _, tc := range tests {
		projects, err := api.FetchAllProjects(tc.host)
		if err != nil {
			t.Fatalf("Got error: %v", err)
		} else if !reflect.DeepEqual(projects, tc.want) {
			t.Fatalf("Got %v;\nwant %v", projects, tc.want)
		}
	}
}
