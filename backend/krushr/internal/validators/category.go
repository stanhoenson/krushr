package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePostCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category is required")
	}

	if category.Position < 0 {
		return errors.New("weight must be a non-negative integer")
	}

	// add more validation rules here

	return nil
}

func ValidatePutCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("category is required")
	}

	if category.Position < 0 {
		return errors.New("weight must be a non-negative integer")
	}

	// add more validation rules here

	return nil
}
