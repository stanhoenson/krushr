package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/utils"
)

func ValidatePostRouteBody(postRouteBody *models.PostRouteBody) error {
	var poiNames []string

	for _, poi := range postRouteBody.PointsOfInterest {
		poiNames = append(poiNames, poi.Name)
		err := ValidatePostPointOfInterestBody(&poi)
		if err != nil {
			return err
		}
	}

	duplicates := utils.FindDuplicates(poiNames)

	if len(duplicates) != 0 {
		return errors.New("duplicate point of interest names")
	}

	return nil
}
