package components

import (
	"fmt"
	"time"
)

var Now = func() time.Time {
	return time.Now()
}

// Copyright generates a copyright string with the current year with a given name.
func Copyright(name string) string {
	year := Now().Year()
	return fmt.Sprintf("© %s %d", name, year)
}
