package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePostCategory(category *models.Category) error {
	if category.Category == "" {
		return errors.New("category is required")
	}

	if category.Icon == "" {
		return errors.New("icon is required")
	}

	if category.Weight < 0 {
		return errors.New("weight must be a non-negative integer")
	}

	if category.TypeID == 0 {
		return errors.New("type_id is required")
	}

	// add more validation rules here

	return nil
}
func ValidatePutCategory(category *models.Category) error {
	if category.Category == "" {
		return errors.New("category is required")
	}

	if category.Icon == "" {
		return errors.New("icon is required")
	}

	if category.Weight < 0 {
		return errors.New("weight must be a non-negative integer")
	}

	if category.TypeID == 0 {
		return errors.New("type_id is required")
	}

	// add more validation rules here

	return nil
}
