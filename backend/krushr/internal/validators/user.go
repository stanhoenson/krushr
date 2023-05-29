package validators

import (
	"regexp"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidateSignUp(user *models.User) error {
	// add more validation rules here

	return nil
}

func ValidatePostUser(user *models.User) error {
	// add more validation rules here

	return nil
}

func ValidatePutUser(user *models.User) error {
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

func isValidPassword(password string) bool {
	return true
}
