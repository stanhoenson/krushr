package services

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/utils"
)

func DeleteRouteByIDAndAuthenticatedUser(ID uint, authenticatedUser *models.User) (*models.Route, error) {
	if authenticatedUser.Role.Name == constants.AdminRoleName {
		return repositories.DeleteEntityByID[models.Route](ID, database.Db)
	}
	return repositories.DeleteRouteByIDAndUserID(ID, authenticatedUser.ID, database.Db)
}

func UpdateRoute(putRouteBody *models.PutRouteBody, authenticatedUser *models.User) (*models.Route, error) {

	pointsOfInterest, err := GetEntitiesByIDs[models.PointOfInterest](&putRouteBody.PointOfInterestIDs)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving points of interest")
	}
	images, err := GetEntitiesByIDs[models.Image](&putRouteBody.ImageIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving images")
	}
	details, err := GetEntitiesByIDs[models.Detail](&putRouteBody.DetailIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving details")
	}
	links, err := GetEntitiesByIDs[models.Link](&putRouteBody.LinkIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving links")
	}
	tx := database.Db.Begin()

	route, err := repositories.GetRouteByIDAndUserID(putRouteBody.ID, authenticatedUser.ID, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	route.Name = putRouteBody.Name
	route.StatusID = putRouteBody.StatusID
	route.Distance =
		utils.PointsOfInterestToDistance(pointsOfInterest)

	updateRoute, err := repositories.UpdateEntity(route, tx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Model(updateRoute).Association("PointsOfInterest").Replace(pointsOfInterest)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(updateRoute).Association("Images").Replace(images)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(updateRoute).Association("Details").Replace(details)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(updateRoute).Association("Links").Replace(links)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}

	tx.Commit()

	return updateRoute, nil

}
func CreateRoute(postRouteBody *models.PostRouteBody, authenticatedUser *models.User) (*models.Route, error) {

	pointsOfInterest, err := GetEntitiesByIDs[models.PointOfInterest](&postRouteBody.PointOfInterestIDs)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving points of interest")
	}
	images, err := GetEntitiesByIDs[models.Image](&postRouteBody.ImageIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving images")
	}
	details, err := GetEntitiesByIDs[models.Detail](&postRouteBody.DetailIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving details")
	}
	links, err := GetEntitiesByIDs[models.Link](&postRouteBody.LinkIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving links")
	}
	tx := database.Db.Begin()

	route := models.Route{
		Name:     postRouteBody.Name,
		StatusID: postRouteBody.StatusID,
		Distance: utils.PointsOfInterestToDistance(pointsOfInterest),
		UserID:   authenticatedUser.ID,
	}
	createdRoute, err := repositories.CreateEntity(&route, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Model(createdRoute).Association("PointsOfInterest").Replace(pointsOfInterest)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(createdRoute).Association("Images").Replace(images)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(createdRoute).Association("Details").Replace(details)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(createdRoute).Association("Links").Replace(links)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}

	tx.Commit()

	return createdRoute, nil

}
