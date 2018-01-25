package user

import (
	"strings"
)

// User struct
type User struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (user *User) validate(tipo string) bool {
	var isValid = true

	if tipo == "insert" {
		if strings.TrimSpace(user.Name) == "" {
			isValid = false
		}
	}

	if tipo == "update" {
		if user.ID == 0 {
			isValid = false
		}

		if strings.TrimSpace(user.Name) == "" {
			isValid = false
		}
	}

	return isValid
}
