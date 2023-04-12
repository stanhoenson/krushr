package validators

import (
	"errors"
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func ValidatePutRoute(route *models.Route) error {
	if route.Title == "" {
		return errors.New("title is required")
	}

	if len(route.Title) > constants.TitleMaxLength {
		return fmt.Errorf("title shouldn't be longer than %d characters", constants.TitleMaxLength)
	}

	statuses, err := repositories.GetEntities[models.Status]()
	if err != nil {
		return errors.New("failed retrieving statuses")
	}
	if route.StatusID < uint(len(*statuses)) {
		return errors.New("status_id should have entry in statuses table")
	}

	users, err := repositories.GetEntities[models.User]()
	if err != nil {
		return errors.New("failed retrieving users")
	}
	if route.StatusID < uint(len(*users)) {
		return errors.New("user_id should have entry in statuses table")
	}

	if len(route.PointsOfInterest) < 2 {
		return errors.New("at least two points of interest is required")
	}

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
