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

func UpdateRoute(ID uint, putRouteBody *models.PutRouteBody, authenticatedUser *models.User, tx *gorm.DB) (*models.Route, error) {
	//create points of interest
	createdOrUpdatedPointsOfInterest := []*models.PointOfInterest{}
	for _, postPointOfInterestBody := range putRouteBody.PointsOfInterest {

		createdPointOfInterest, err := FindOrCreateOrUpdatePointOfInterest(&postPointOfInterestBody, authenticatedUser, tx)
		if err != nil {
			return nil, err
		}
		createdOrUpdatedPointsOfInterest = append(createdOrUpdatedPointsOfInterest, createdPointOfInterest)
	}

	//find or create categories
	foundOrCreatedCategories := []*models.Category{}
	for _, postCategoryBody := range putRouteBody.Categories {

		foundOrCreatedCategory, err := FirstOrCreateCategory(&postCategoryBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedCategories = append(foundOrCreatedCategories, foundOrCreatedCategory)
	}

	//TODO to error or not to error on empty result
	images, err := GetEntitiesByIDs[models.Image](&putRouteBody.ImageIDs)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving images")
	}

	//create links
	foundOrCreatedLinks := []*models.Link{}
	for _, postLinkBody := range putRouteBody.Links {

		foundOrCreatedLink, err := FirstOrCreateLink(&postLinkBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedLinks = append(foundOrCreatedLinks, foundOrCreatedLink)
	}

	//create details
	foundOrCreatedDetails := []*models.Detail{}
	for _, postDetailBody := range putRouteBody.Details {

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

	status, err := repositories.GetEntityByID[models.Status](putRouteBody.StatusID, tx)
	if err != nil {
		return nil, err
	}

	route, err := repositories.GetRouteByIDAndUserID(ID, authenticatedUser.ID, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	route.Name = putRouteBody.Name
	route.Status = *status
	route.Distance = utils.PointsOfInterestToDistance(createdOrUpdatedPointsOfInterest)
	route.User = *authenticatedUser
	route.PointsOfInterest = createdOrUpdatedPointsOfInterest
	route.Links = foundOrCreatedLinks
	route.Details = foundOrCreatedDetails
	route.Images = imagesPointers

	updatedRoute, err := repositories.UpdateEntity(route, tx)
	if err != nil {
		return nil, err
	}

	for index, pointOfInterest := range route.PointsOfInterest {

		routePointOfInterest := models.RoutesPointsOfInterest{
			RouteID:           route.ID,
			PointOfInterestID: pointOfInterest.ID,
			Position:          uint(index),
		}
		_, err := repositories.UpdateEntity(&routePointOfInterest, tx)
		if err != nil {
			return nil, err
		}
	}

	return updatedRoute, nil
}

func CreateRoute(postRouteBody *models.PostRouteBody, authenticatedUser *models.User, tx *gorm.DB) (*models.Route, error) {

	//create points of interest
	createdOrUpdatedPointsOfInterest := []*models.PointOfInterest{}
	for _, postPointOfInterestBody := range postRouteBody.PointsOfInterest {

		createdPointOfInterest, err := FindOrCreateOrUpdatePointOfInterest(&postPointOfInterestBody, authenticatedUser, tx)
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

	status, err := repositories.GetEntityByID[models.Status](postRouteBody.StatusID, tx)
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
	for index, pointOfInterest := range route.PointsOfInterest {

		routePointOfInterest := models.RoutesPointsOfInterest{
			RouteID:           route.ID,
			PointOfInterestID: pointOfInterest.ID,
			Position:          uint(index),
		}
		_, err := repositories.UpdateEntity(&routePointOfInterest, tx)
		if err != nil {
			return nil, err
		}
	}

	return createdRoute, nil

}
