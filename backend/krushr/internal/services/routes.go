package services

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/utils"
	"gorm.io/gorm"
)

func DeleteRouteByIDAndAuthenticatedUser(ID uint, authenticatedUser *models.User) (*models.Route, error) {
	if authenticatedUser.Role.Name == constants.AdminRoleName {
		return repositories.DeleteEntityByID[models.Route](ID, database.Db)
	}
	return repositories.DeleteRouteByIDAndUserID(ID, authenticatedUser.ID, database.Db)
}

func UpdateRoute(putRouteBody *models.PutRouteBody, authenticatedUser *models.User) (*models.Route, error) {
	// pointsOfInterest, err := GetEntitiesByIDs[models.PointOfInterest](&postRouteBody.PointOfInterestIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving points of interest")
	// }
	// images, err := GetEntitiesByIDs[models.Image](&postRouteBody.ImageIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving images")
	// }
	// details, err := GetEntitiesByIDs[models.Detail](&postRouteBody.DetailIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving details")
	// }
	// links, err := GetEntitiesByIDs[models.Link](&postRouteBody.LinkIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving links")
	// }
	tx := database.Db.Begin()

	//create points of interest
	createdOrUpdatedPointsOfInterest := []*models.PointOfInterest{}
	for _, postPointOfInterestBody := range putRouteBody.PointsOfInterest {

		createdPointOfInterest, err := CreateOrUpdatePointOfInterest(&postPointOfInterestBody, authenticatedUser, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		createdOrUpdatedPointsOfInterest = append(createdOrUpdatedPointsOfInterest, createdPointOfInterest)
	}
	images, err := GetEntitiesByIDs[models.Image](&putRouteBody.ImageIDs)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error retrieving images")
	}

	//create links
	foundOrCreatedLinks := []*models.Link{}
	for _, postLinkBody := range putRouteBody.Links {

		foundOrCreatedLink, err := FirstOrCreateLink(&postLinkBody, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foundOrCreatedLinks = append(foundOrCreatedLinks, foundOrCreatedLink)
	}

	//create details
	foundOrCreatedDetails := []*models.Detail{}
	for _, postDetailBody := range putRouteBody.Details {

		foundOrCreatedDetail, err := FirstOrCreateDetail(&postDetailBody, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		foundOrCreatedDetails = append(foundOrCreatedDetails, foundOrCreatedDetail)
	}

	imagesPointers := []*models.Image{}
	for _, image := range *images {
		imagesPointers = append(imagesPointers, &image)
	}

	route, err := repositories.GetRouteByIDAndUserID(putRouteBody.ID, authenticatedUser.ID, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	route.Name = putRouteBody.Name
	route.StatusID = putRouteBody.StatusID
	route.Distance = utils.PointsOfInterestToDistance(createdOrUpdatedPointsOfInterest)
	route.UserID = authenticatedUser.ID
	route.PointsOfInterest = createdOrUpdatedPointsOfInterest
	route.Links = foundOrCreatedLinks
	route.Details = foundOrCreatedDetails
	route.Images = imagesPointers

	updatedRoute, err := repositories.UpdateEntity(route, tx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// err = tx.Model(createdRoute).Association("PointsOfInterest").Replace(pointsOfInterest)
	for index, pointOfInterest := range route.PointsOfInterest {
		var routePointOfInterest models.RoutesPointsOfInterest
		//TODO should this be abstracted away to a repo function?
		result := tx.Where("route_id = ?", route.ID).Where("point_of_interest_id = ?", pointOfInterest.ID).First(&routePointOfInterest)
		if result.Error != nil {
			tx.Rollback()
			return nil, err
		}
		routePointOfInterest.Position = uint(index)
		result = tx.Updates(&routePointOfInterest)
		if result.Error != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	// err = tx.Model(createdRoute).Association("Images").Replace(images)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx.Model(createdRoute).Association("Details").Replace(details)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx.Model(createdRoute).Association("Links").Replace(links)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }

	tx.Commit()

	return updatedRoute, nil
}

func CreateRoute(postRouteBody *models.PostRouteBody, authenticatedUser *models.User, tx *gorm.DB) (*models.Route, error) {

	// pointsOfInterest, err := GetEntitiesByIDs[models.PointOfInterest](&postRouteBody.PointOfInterestIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving points of interest")
	// }
	// images, err := GetEntitiesByIDs[models.Image](&postRouteBody.ImageIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving images")
	// }
	// details, err := GetEntitiesByIDs[models.Detail](&postRouteBody.DetailIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving details")
	// }
	// links, err := GetEntitiesByIDs[models.Link](&postRouteBody.LinkIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving links")
	// }

	//create points of interest
	createdOrUpdatedPointsOfInterest := []*models.PointOfInterest{}
	for _, postPointOfInterestBody := range postRouteBody.PointsOfInterest {

		createdPointOfInterest, err := CreateOrUpdatePointOfInterest(&postPointOfInterestBody, authenticatedUser, tx)
		if err != nil {
			return nil, err
		}
		createdOrUpdatedPointsOfInterest = append(createdOrUpdatedPointsOfInterest, createdPointOfInterest)
	}

	//find or create categories
	foundOrCreatedCategories := []*models.Category{}
	for _, postCategoryBody := range postRouteBody.Categories {

		foundOrCreatedCategory, err := FirstOrCreateCategory(&postCategoryBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedCategories = append(foundOrCreatedCategories, foundOrCreatedCategory)
	}

	images, err := GetEntitiesByIDs[models.Image](&postRouteBody.ImageIDs)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving images")
	}

	//create links
	foundOrCreatedLinks := []*models.Link{}
	for _, postLinkBody := range postRouteBody.Links {

		foundOrCreatedLink, err := FirstOrCreateLink(&postLinkBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedLinks = append(foundOrCreatedLinks, foundOrCreatedLink)
	}

	//create details
	foundOrCreatedDetails := []*models.Detail{}
	for _, postDetailBody := range postRouteBody.Details {

		foundOrCreatedDetail, err := FirstOrCreateDetail(&postDetailBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedDetails = append(foundOrCreatedDetails, foundOrCreatedDetail)
	}

	imagesPointers := []*models.Image{}
	for _, image := range *images {
		imagesPointers = append(imagesPointers, &image)
	}

	status, err := repositories.GetEntity[models.Status](postRouteBody.StatusID, tx)
	if err != nil {
		return nil, err
	}

	route := models.Route{
		Name:             postRouteBody.Name,
		Status:           *status,
		Distance:         utils.PointsOfInterestToDistance(createdOrUpdatedPointsOfInterest),
		User:             *authenticatedUser,
		PointsOfInterest: createdOrUpdatedPointsOfInterest,
		Links:            foundOrCreatedLinks,
		Details:          foundOrCreatedDetails,
		Images:           imagesPointers,
	}
	createdRoute, err := repositories.CreateEntity(&route, tx)
	if err != nil {
		return nil, err
	}
	// err = tx.Model(createdRoute).Association("PointsOfInterest").Replace(pointsOfInterest)
	for index, pointOfInterest := range route.PointsOfInterest {
		var routePointOfInterest models.RoutesPointsOfInterest
		//TODO should this be abstracted away to a repo function?
		result := tx.Where("route_id = ?", route.ID).Where("point_of_interest_id = ?", pointOfInterest.ID).First(&routePointOfInterest)
		if result.Error != nil {
			return nil, err
		}
		routePointOfInterest.Position = uint(index)
		result = tx.Updates(&routePointOfInterest)
		if result.Error != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, fmt.Errorf("Error replacing association")
	}
	// err = tx.Model(createdRoute).Association("Images").Replace(images)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx.Model(createdRoute).Association("Details").Replace(details)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx.Model(createdRoute).Association("Links").Replace(links)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }

	return createdRoute, nil

}
