package services

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetRouteByIDWithAssociations(ID uint) (*models.Route, error) {
	route, err := repositories.GetEntityByIDWithAssociations[models.Route](ID, clause.Associations, database.Db)
	if err != nil {
		return nil, err
	}
	pointsOfInterest, err := GetPointsOfInterestByRouteIDOrderedByPositionWithAssociations(route.ID)
	if err != nil {
		return nil, err
	}

	for index := range route.PointsOfInterest {
		route.PointsOfInterest[index] = &(*pointsOfInterest)[index]
	}

	return route, nil
}

func GetRoutesWithAssociationsByUserID(userID uint) (*[]models.Route, error) {
	return repositories.GetRoutesWithAssociationsByUserID(userID, database.Db)
}

func GetPublishedRoutes() (*[]models.Route, error) {
	return repositories.GetPublishedRoutes(database.Db)
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
	tx.Model(&route).Association("PointsOfInterest").Replace(route.PointsOfInterest)

	replaceRouteLinkAssociations(route, &routeRelatedEntities.links, tx)

	replaceRouteDetailAssociations(route, &routeRelatedEntities.details, tx)

	replaceRouteImageAssociations(route, &routeRelatedEntities.images, tx)

	route.Categories = routeRelatedEntities.categories
	tx.Model(&route).Association("Categories").Replace(route.Categories)

	updatedRoute, err := repositories.UpdateEntity(route, tx)
	if err != nil {
		return nil, err
	}

	for index, pointOfInterest := range updatedRoute.PointsOfInterest {

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

func replaceRouteImageAssociations(route *models.Route, images *[]*models.Image, tx *gorm.DB) error {
	var oldImages []models.Image
	for _, image := range route.Images {
		oldImages = append(oldImages, *image)
	}
	tx.Model(&route).Association("Images").Delete(route.Images)
	tx.Model(&route).Association("Images").Append(images)
	for _, image := range oldImages {
		var count int64
		//TODO raw sql might just be better
		tx.Raw("SELECT COUNT(*) FROM routes_images WHERE image_id = ?", image.ID).Scan(&count)
		if count == 0 {
			tx.Delete(image, image)
		}

	}
	route.Images = *images
	return nil

}
func replaceRouteLinkAssociations(route *models.Route, links *[]*models.Link, tx *gorm.DB) error {
	var oldLinks []models.Link
	for _, link := range route.Links {
		oldLinks = append(oldLinks, *link)
	}
	tx.Model(&route).Association("Links").Delete(route.Images)
	tx.Model(&route).Association("Links").Append(links)
	for _, link := range oldLinks {
		var count int64
		//TODO raw sql might just be better
		tx.Raw("SELECT COUNT(*) FROM routes_links WHERE link_id = ?", link.ID).Scan(&count)
		if count == 0 {
			tx.Delete(link, link)
		}

	}
	route.Links = *links
	return nil

}

func replaceRouteDetailAssociations(route *models.Route, details *[]*models.Detail, tx *gorm.DB) error {
	var oldDetails []models.Detail
	for _, detail := range route.Details {
		oldDetails = append(oldDetails, *detail)
	}
	tx.Model(&route).Association("Details").Delete(route.Details)
	tx.Model(&route).Association("Details").Append(details)
	for _, detail := range oldDetails {
		var count int64
		//TODO raw sql might just be better
		tx.Raw("SELECT COUNT(*) FROM routes_details WHERE detail_id = ?", detail.ID).Scan(&count)
		if count == 0 {
			tx.Delete(detail, detail)
		}

	}
	route.Details = *details
	return nil

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
	fmt.Println(images)
	if err != nil {
		return nil, fmt.Errorf("error retrieving images")
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
		img := image
		imagesPointers = append(imagesPointers, &img)
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
		Categories:       routeRelatedEntities.categories,
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
