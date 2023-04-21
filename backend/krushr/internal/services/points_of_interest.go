package services

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"gorm.io/gorm"
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
	// details, err := GetEntitiesByIDs[models.Detail](&putPointOfInterestBody.DetailIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving details")
	// }
	// links, err := GetEntitiesByIDs[models.Link](&putPointOfInterestBody.LinkIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving links")
	// }
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
	// err = tx.Model(updatedPointOfInterest).Association("Details").Replace(details)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx.Model(updatedPointOfInterest).Association("Links").Replace(links)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }

	tx.Commit()

	return updatedPointOfInterest, nil
}
func CreateOrUpdatePointOfInterest(postPointOfInterestBody *models.PostPointOfInterestBody, authenticatedUser *models.User, tx *gorm.DB) (*models.PointOfInterest, error) {

	images, err := repositories.GetEntitiesByIDs[models.Image](&postPointOfInterestBody.ImageIDs, tx)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving images" + err.Error())
	}

	//find or create links
	foundOrCreatedLinks := []*models.Link{}
	for _, postLinkBody := range postPointOfInterestBody.Links {

		foundOrCreatedLink, err := FirstOrCreateLink(&postLinkBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedLinks = append(foundOrCreatedLinks, foundOrCreatedLink)
	}

	//find or create details
	foundOrCreatedDetails := []*models.Detail{}
	for _, postDetailBody := range postPointOfInterestBody.Details {

		foundOrCreatedDetail, err := FirstOrCreateDetail(&postDetailBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedDetails = append(foundOrCreatedDetails, foundOrCreatedDetail)
	}

	//find or create categories
	foundOrCreatedCategories := []*models.Category{}
	for _, postCategoryBody := range postPointOfInterestBody.Categories {

		foundOrCreatedCategory, err := FirstOrCreateCategory(&postCategoryBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedCategories = append(foundOrCreatedCategories, foundOrCreatedCategory)
	}

	imagesPointers := []*models.Image{}
	for _, image := range *images {
		imagesPointers = append(imagesPointers, &image)
	}

	pointOfInterest := models.PointOfInterest{
		Name:       postPointOfInterestBody.Name,
		Latitude:   postPointOfInterestBody.Latitude,
		Longitude:  postPointOfInterestBody.Longitude,
		Categories: foundOrCreatedCategories,
		UserID:     authenticatedUser.ID,
		Details:    foundOrCreatedDetails,
		Links:      foundOrCreatedLinks,
		Images:     imagesPointers,
	}

	createdPointOfInterest, err := repositories.FirstOrCreateEntity(&pointOfInterest, &pointOfInterest, tx)
	// if err != nil {
	// 	tx2.Rollback()
	// 	return nil, err
	// }
	// err = tx2.Model(createdPointOfInterest).Association("Images").Replace(images)
	// if err != nil {
	// 	tx2.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx2.Model(createdPointOfInterest).Association("Details").Replace(details)
	// if err != nil {
	// 	tx2.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx2.Model(createdPointOfInterest).Association("Links").Replace(links)
	// if err != nil {
	// 	tx2.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }

	return createdPointOfInterest, nil
}

func CreatePointOfInterest(postPointOfInterestBody *models.PostPointOfInterestBody, authenticatedUser *models.User) (*models.PointOfInterest, error) {
	images, err := GetEntitiesByIDs[models.Image](&postPointOfInterestBody.ImageIDs)
	if err != nil {
		return nil, fmt.Errorf("Error retrieving images")
	}
	// details, err := GetEntitiesByIDs[models.Detail](&postPointOfInterestBody.DetailIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving details")
	// }
	// links, err := GetEntitiesByIDs[models.Link](&postPointOfInterestBody.LinkIDs)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error retrieving links")
	// }
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
	// err = tx.Model(createdPointOfInterest).Association("Details").Replace(details)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }
	// err = tx.Model(createdPointOfInterest).Association("Links").Replace(links)
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, fmt.Errorf("Error replacing association")
	// }

	tx.Commit()

	return createdPointOfInterest, nil
}
