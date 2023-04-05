package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePostEntry(entry *models.Entry) error {
	if entry.Content == "" {
		return errors.New("content is required")
	}

	if entry.Hyperlink == "" {
		return errors.New("hyperlink is required")
	}

	if entry.TypeID == 0 {
		return errors.New("type_id is required")
	}

	// add more validation rules here

	return nil
}
