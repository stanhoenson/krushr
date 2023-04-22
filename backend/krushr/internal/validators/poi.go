package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePutPointOfInterest(poi *models.PutPointOfInterestBody) error {
	if poi.Name == "" {
		return errors.New("title is required")
	}

	if poi.Longitude < -180 || poi.Longitude > 180 {
		return errors.New("invalid longitude value")
	}

	if poi.Latitude < -90 || poi.Latitude > 90 {
		return errors.New("invalid latitude value")
	}

	if len(poi.Categories) == 0 {
		return errors.New("at least one category is required")
	}

	// add more validation rules here

	return nil
}

func ValidatePostPointOfInterest(poi *models.PostPointOfInterestBody) error {
	if poi.Name == "" {
		return errors.New("title is required")
	}

	if poi.Longitude < -180 || poi.Longitude > 180 {
		return errors.New("invalid longitude value")
	}

	if poi.Latitude < -90 || poi.Latitude > 90 {
		return errors.New("invalid latitude value")
	}

	if len(poi.Categories) == 0 {
		return errors.New("at least one category is required")
	}

	// add more validation rules here

	return nil
}
