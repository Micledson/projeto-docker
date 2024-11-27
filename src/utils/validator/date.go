package validator

import (
	"regexp"
	"time"
)

func IsDateValid(date *time.Time) bool {
	if date == nil || date.IsZero() {
		return false
	}
	if date.After(time.Now()) {
		return false
	}
	r := regexp.MustCompile("^[0-9]{4}-?[0-9]{2}-?[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2} \\+\\d{4} UTC$")
	return r.MatchString(date.String())
}
