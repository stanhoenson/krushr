package validators

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
)

func ValidatePostPointOfInterestBody(postPointOfInterestBody *models.PostPointOfInterestBody) error {

	if !postPointOfInterestBody.Support {
		if len(postPointOfInterestBody.Details) < 1 {
			return errors.New("point of interest doesn't have a Detail")
		}

		for _, detail := range postPointOfInterestBody.Details {
			err := ValidateStruct(detail)
			if err != nil {
				return err
			}
		}
		for _, link := range postPointOfInterestBody.Links {
			err := ValidateStruct(link)
			if err != nil {
				return err
			}
		}
		if len(postPointOfInterestBody.ImageIDs) < 1 {
			return errors.New("point of interest doesn't have a ImageID")
		}
		for _, categories := range postPointOfInterestBody.Categories {
			err := ValidateStruct(categories)
			if err != nil {
				return err
			}
		}

	}

	return nil

}
