package utils

import (
	"github.com/thalissonfelipe/ugly-api/models"
)

func ValidateCreateUserBody(user models.User) error {
	if user.ID == "" || user.Name == "" || user.Username == "" || user.Password == "" {
		return ErrMissingRequiredFields
	}

	if len(user.Username) < 6 || len(user.Username) > 20 {
		return ErrInvalidUsernameLength
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		return ErrInvalidPasswordLength
	}

	return nil
}
