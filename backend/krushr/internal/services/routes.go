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

func GetRouteByIDWithAssociations(ID uint) (*models.Route, error) {
	return repositories.GetRouteByIDWithAssociations(ID, database.Db)
}

func DeleteRouteByIDAndAuthenticatedUser(ID uint, authenticatedUser *models.User) (*models.Route, error) {
	if authenticatedUser.Role.Name == constants.AdminRoleName {
		return repositories.DeleteEntityByID[models.Route](ID, database.Db)
	}
	return repositories.DeleteRouteByIDAndUserID(ID, authenticatedUser.ID, database.Db)
}

func UpdateRoute(ID uint, putRouteBody *models.PutRouteBody, authenticatedUser *models.User, tx *gorm.DB) (*models.Route, error) {
	routeRelatedEntities, err := CreateOrUpdateRouteRelatedEntities(&putRouteBody.PostRouteBody, authenticatedUser, tx)
	if err != nil {
		return nil, err
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
	route.Distance = utils.PointsOfInterestToDistance(routeRelatedEntities.pointsOfInterest)
	route.User = *authenticatedUser
	route.PointsOfInterest = routeRelatedEntities.pointsOfInterest
	route.Links = routeRelatedEntities.links
	route.Details = routeRelatedEntities.details
	route.Images = routeRelatedEntities.images

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

type routeRelatedEntities struct {
	pointsOfInterest []*models.PointOfInterest
	images           []*models.Image
	links            []*models.Link
	details          []*models.Detail
	categories       []*models.Category
}

func CreateOrUpdateRouteRelatedEntities(postRouteBody *models.PostRouteBody, authenticatedUser *models.User, tx *gorm.DB) (*routeRelatedEntities, error) {
	// create points of interest
	createdOrUpdatedPointsOfInterest := []*models.PointOfInterest{}
	for _, postPointOfInterestBody := range postRouteBody.PointsOfInterest {

		createdPointOfInterest, err := FindOrCreateOrUpdatePointOfInterest(&postPointOfInterestBody, authenticatedUser, tx)
		if err != nil {
			return nil, err
		}
		createdOrUpdatedPointsOfInterest = append(createdOrUpdatedPointsOfInterest, createdPointOfInterest)
	}

	// find or create categories
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

	// create links
	foundOrCreatedLinks := []*models.Link{}
	for _, postLinkBody := range postRouteBody.Links {

		foundOrCreatedLink, err := FirstOrCreateLink(&postLinkBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedLinks = append(foundOrCreatedLinks, foundOrCreatedLink)
	}

	// create details
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

	return &routeRelatedEntities{
		images:           imagesPointers,
		details:          foundOrCreatedDetails,
		categories:       foundOrCreatedCategories,
		links:            foundOrCreatedLinks,
		pointsOfInterest: createdOrUpdatedPointsOfInterest,
	}, nil
}

func CreateRoute(postRouteBody *models.PostRouteBody, authenticatedUser *models.User, tx *gorm.DB) (*models.Route, error) {
	routeRelatedEntities, err := CreateOrUpdateRouteRelatedEntities(postRouteBody, authenticatedUser, tx)
	if err != nil {
		return nil, err
	}

	status, err := repositories.GetEntityByID[models.Status](postRouteBody.StatusID, tx)
	if err != nil {
		return nil, err
	}

	route := models.Route{
		Name:             postRouteBody.Name,
		Status:           *status,
		Distance:         utils.PointsOfInterestToDistance(routeRelatedEntities.pointsOfInterest),
		User:             *authenticatedUser,
		PointsOfInterest: routeRelatedEntities.pointsOfInterest,
		Links:            routeRelatedEntities.links,
		Details:          routeRelatedEntities.details,
		Images:           routeRelatedEntities.images,
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
