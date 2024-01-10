package utils

import (
	"fmt"
	"time"
)

func Copyright(name string) string {
	year := Now().Year()
	return fmt.Sprintf("© %s %d", name, year)
}

// Functions to manipulate runtime behaviour for tests

var Fetch = func(path string) (string, error) {
	return FetchURL(path)
}

var Now = func() time.Time {
	return time.Now()
}
