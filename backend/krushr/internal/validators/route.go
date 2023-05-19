package validators

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePutRoute(route *models.PutRouteBody) error {

	for _, poi := range route.PointsOfInterest {
		if !poi.Support && len(poi.ImageIDs) == 0 {
			return fmt.Errorf("point of interest has no image")
		}
	}

	return nil
}

func ValidatePostRouteBody(route *models.PostRouteBody) error {

	for _, poi := range route.PointsOfInterest {
		if !poi.Support && len(poi.ImageIDs) == 0 {
			return fmt.Errorf("point of interest has no image")
		}
	}
	return nil
}
