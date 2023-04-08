package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePutRoute(route *models.Route) error {
	if route.Title == "" {
		return errors.New("title is required")
	}

	if route.StatusID == 0 {
		return errors.New("status_id is required")
	}

	if route.UserID == 0 {
		return errors.New("user_id is required")
	}

	if len(route.PointsOfInterest) == 0 {
		return errors.New("at least one point of interest is required")
	}

	// add more validation rules here

	return nil
}

func ValidatePostRoute(route *models.Route) error {
	if route.Title == "" {
		return errors.New("title is required")
	}

	if route.StatusID == 0 {
		return errors.New("status_id is required")
	}

	if route.UserID == 0 {
		return errors.New("user_id is required")
	}

	if len(route.PointsOfInterest) == 0 {
		return errors.New("at least one point of interest is required")
	}

	// add more validation rules here

	return nil
}
