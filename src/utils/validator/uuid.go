package validator

import (
	"github.com/google/uuid"
	"regexp"
)

func IsUUIDValid(id uuid.UUID) bool {
	if id == uuid.Nil || id.ID() == 0 {
		return false
	}

	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(id.String())
}
