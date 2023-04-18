package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePostCategoryBody(postCategoryBody *models.PostCategoryBody) error {
	if postCategoryBody.Name == "" {
		return errors.New("category is required")
	}

	if postCategoryBody.Position < 0 {
		return errors.New("weight must be a non-negative integer")
	}

	// add more validation rules here

	return nil
}

func ValidatePutCategoryBody(putCategoryBody *models.PutCategoryBody) error {
	if putCategoryBody.Name == "" {
		return errors.New("category is required")
	}

	if putCategoryBody.Position < 0 {
		return errors.New("weight must be a non-negative integer")
	}

	// add more validation rules here

	return nil
}
