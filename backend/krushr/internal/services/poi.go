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

func UpdatePointOfInterest(ID uint, putPointOfInterestBody *models.PutPointOfInterestBody, authenticatedUser *models.User, tx *gorm.DB) (*models.PointOfInterest, error) {
	//find

	retrievedPointOfInterest, err := repositories.GetPointOfInterestByIDAndUserID(ID, authenticatedUser.ID, tx)

	if err != nil {
		return nil, err
	}

	pointOfInterestRelatedEntries, err := CreateOrUpdateRelatedEntities(&putPointOfInterestBody.PostPointOfInterestBody, tx)

	if err != nil {
		return nil, err
	}

	retrievedPointOfInterest.Name = putPointOfInterestBody.Name
	retrievedPointOfInterest.Longitude = putPointOfInterestBody.Longitude
	retrievedPointOfInterest.Latitude = putPointOfInterestBody.Latitude
	retrievedPointOfInterest.Categories = pointOfInterestRelatedEntries.categories
	retrievedPointOfInterest.Details = pointOfInterestRelatedEntries.details
	retrievedPointOfInterest.Links = pointOfInterestRelatedEntries.links
	retrievedPointOfInterest.Images = pointOfInterestRelatedEntries.images

	updatePointOfInterest, err := repositories.UpdateEntity(retrievedPointOfInterest, tx)
	if err != nil {
		return nil, err
	}

	return updatePointOfInterest, err
}

type PointOfInterestRelatedEntries struct {
	images     []*models.Image
	links      []*models.Link
	details    []*models.Detail
	categories []*models.Category
}

func CreateOrUpdateRelatedEntities(postPointOfInterestBody *models.PostPointOfInterestBody, tx *gorm.DB) (*PointOfInterestRelatedEntries, error) {

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

	return &PointOfInterestRelatedEntries{
		images:     imagesPointers,
		categories: foundOrCreatedCategories,
		details:    foundOrCreatedDetails,
		links:      foundOrCreatedLinks,
	}, nil

}

func FindOrCreateOrUpdatePointOfInterest(postPointOfInterestBody *models.PostPointOfInterestBody, authenticatedUser *models.User, tx *gorm.DB) (*models.PointOfInterest, error) {

	//find
	getPointOfInterestBody := postPointOfInterestBody.ToGetPointOfInterestBody()

	pointOfInterest := getPointOfInterestBody.ToPointOfInterest()

	retrievedPointOfInterest, err := repositories.GetEntityByConditions(&pointOfInterest, tx)

	if err != nil {
		return nil, err
	}

	if retrievedPointOfInterest.ID != 0 {
		//create

		pointOfInterestRelatedEntries, err := CreateOrUpdateRelatedEntities(postPointOfInterestBody, tx)

		if err != nil {
			return nil, err
		}

		pointOfInterest := models.PointOfInterest{
			Name:       postPointOfInterestBody.Name,
			Longitude:  postPointOfInterestBody.Longitude,
			Latitude:   postPointOfInterestBody.Latitude,
			Categories: pointOfInterestRelatedEntries.categories,
			UserID:     authenticatedUser.ID,
			Details:    pointOfInterestRelatedEntries.details,
			Links:      pointOfInterestRelatedEntries.links,
			Images:     pointOfInterestRelatedEntries.images,
		}

		createdPointOfInterest, err := repositories.CreateEntity(&pointOfInterest, tx)

		if err != nil {
			return nil, err
		}

		return createdPointOfInterest, err
	} else if retrievedPointOfInterest.UserID == authenticatedUser.ID {
		//update

		pointOfInterestRelatedEntries, err := CreateOrUpdateRelatedEntities(postPointOfInterestBody, tx)

		if err != nil {
			return nil, err
		}

		retrievedPointOfInterest.Name = postPointOfInterestBody.Name
		retrievedPointOfInterest.Longitude = postPointOfInterestBody.Longitude
		retrievedPointOfInterest.Latitude = postPointOfInterestBody.Latitude
		retrievedPointOfInterest.Categories = pointOfInterestRelatedEntries.categories
		retrievedPointOfInterest.Details = pointOfInterestRelatedEntries.details
		retrievedPointOfInterest.Links = pointOfInterestRelatedEntries.links
		retrievedPointOfInterest.Images = pointOfInterestRelatedEntries.images

		updatePointOfInterest, err := repositories.UpdateEntity(retrievedPointOfInterest, tx)
		if err != nil {
			return nil, err
		}

		return updatePointOfInterest, err
	} else {
		return retrievedPointOfInterest, nil
	}

	// createdPointOfInterest, err := repositories.FirstOrCreateEntity(&pointOfInterest, &pointOfInterest, tx)
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
