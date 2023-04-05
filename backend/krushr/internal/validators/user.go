package validators

import (
	"errors"
	"regexp"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePostUser(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if user.RoleID == 0 {
		return errors.New("role_id is required")
	}

	// add more validation rules here

	return nil
}
func ValidatePutUser(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if user.RoleID == 0 {
		return errors.New("role_id is required")
	}

	// add more validation rules here

	return nil
}

// isValidEmail returns true if the email address is valid, false otherwise.
func isValidEmail(email string) bool {
	// This regex pattern is a simplified version of the official email regex pattern.
	// You can use a more strict or more relaxed regex pattern based on your requirements.
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	return err == nil && match
}
