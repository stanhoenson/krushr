package validators

import (
	"errors"
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePutRoute(route *models.PutRouteBody) error {
	if route.Name == "" {
		return errors.New("title is required")
	}

	if len(route.Name) > constants.TitleMaxLength {
		return fmt.Errorf("title shouldn't be longer than %d characters", constants.TitleMaxLength)
	}

	// TODO probably caught by insert
	// if len(route.PointsOfInterest) < 2 {
	// 	return errors.New("at least two points of interest is required")
	// }

	return nil
}

func ValidatePostRouteBody(route *models.PostRouteBody) error {
	if route.Name == "" {
		return errors.New("title is required")
	}

	if route.StatusID == 0 {
		return errors.New("status_id is required")
	}

	return nil
}
