package services

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func DeletePointOfInterestByIDAndAuthentictedUser(ID uint, authenticatedUser *models.User) (*models.Route, error) {
	if authenticatedUser.Role.Name == constants.AdminRoleName {
		return repositories.DeleteEntityByID[models.Route](ID, database.Db)
	}
	return repositories.DeleteRouteByIDAndUserID(ID, authenticatedUser.ID, database.Db)
}

func UpdatePointOfInterest(putPointOfInterestBody *models.PutPointOfInterestBody, authenticatedUser *models.User) (*models.PointOfInterest, error) {

	images, err := GetEntitiesByIDs[models.Image](&putPointOfInterestBody.ImageIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving images")
	}
	details, err := GetEntitiesByIDs[models.Detail](&putPointOfInterestBody.DetailIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving details")
	}
	links, err := GetEntitiesByIDs[models.Link](&putPointOfInterestBody.LinkIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving links")
	}
	tx := database.Db.Begin()

	pointOfInterest, err := repositories.GetPointOfInterestByIDAndUserID(putPointOfInterestBody.ID, authenticatedUser.ID, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	pointOfInterest.Name = putPointOfInterestBody.Name
	pointOfInterest.Latitude = putPointOfInterestBody.Latitude
	pointOfInterest.Longitude = putPointOfInterestBody.Longitude

	updatedPointOfInterest, err := repositories.UpdateEntity(pointOfInterest, tx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Model(updatedPointOfInterest).Association("Images").Replace(images)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(updatedPointOfInterest).Association("Details").Replace(details)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(updatedPointOfInterest).Association("Links").Replace(links)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}

	tx.Commit()

	return updatedPointOfInterest, nil

}

func CreatePointOfInterest(postPointOfInterestBody *models.PostPointOfInterestBody, authenticatedUser *models.User) (*models.PointOfInterest, error) {

	images, err := GetEntitiesByIDs[models.Image](&postPointOfInterestBody.ImageIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving images")
	}
	details, err := GetEntitiesByIDs[models.Detail](&postPointOfInterestBody.DetailIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving details")
	}
	links, err := GetEntitiesByIDs[models.Link](&postPointOfInterestBody.LinkIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving links")
	}
	tx := database.Db.Begin()

	pointOfInterest := models.PointOfInterest{
		Name:      postPointOfInterestBody.Name,
		Latitude:  postPointOfInterestBody.Latitude,
		Longitude: postPointOfInterestBody.Longitude,
		UserID:    authenticatedUser.ID,
	}

	createdPointOfInterest, err := repositories.CreateEntity(&pointOfInterest, tx)

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Model(createdPointOfInterest).Association("Images").Replace(images)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(createdPointOfInterest).Association("Details").Replace(details)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}
	err = tx.Model(createdPointOfInterest).Association("Links").Replace(links)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("Error replacing association")
	}

	tx.Commit()

	return createdPointOfInterest, nil

}
