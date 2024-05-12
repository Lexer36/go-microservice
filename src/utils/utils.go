package utils

import "time"

const (
	layout = "2006-01-02 15:04:05"
)

func ValidateTimeFormat(s string) error {
	_, err := time.Parse(layout, s)
	return err
}
