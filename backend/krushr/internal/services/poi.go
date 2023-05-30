package services

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"gorm.io/gorm"
)

func DeletePointOfInterestByIDAndAuthentictedUser(ID uint, authenticatedUser *models.User) (uint, error) {
	if authenticatedUser.Role.Name == constants.AdminRoleName {
		return repositories.DeleteEntityByID[models.PointOfInterest](ID, database.Db)
	}
	return repositories.DeletePointOfInterestByIDAndAuthentictedUser(ID, authenticatedUser.ID, database.Db)
}

func UpdatePointOfInterest(ID uint, putPointOfInterestBody *models.PutPointOfInterestBody, authenticatedUser *models.User, tx *gorm.DB) (*models.PointOfInterest, error) {
	// find

	retrievedPointOfInterest, err := repositories.GetPointOfInterestByIDAndUserID(ID, authenticatedUser.ID, tx)
	if err != nil {
		return nil, err
	}

	pointOfInterestRelatedEntries, err := CreateOrUpdatePointOfInterestRelatedEntities(putPointOfInterestBody, tx)
	if err != nil {
		return nil, err
	}

	retrievedPointOfInterest.Support = putPointOfInterestBody.Support
	retrievedPointOfInterest.Name = putPointOfInterestBody.Name
	retrievedPointOfInterest.Longitude = putPointOfInterestBody.Longitude
	retrievedPointOfInterest.Latitude = putPointOfInterestBody.Latitude
	retrievedPointOfInterest.Categories = pointOfInterestRelatedEntries.categories
	tx.Model(&retrievedPointOfInterest).Association("Categories").Replace(retrievedPointOfInterest.Categories)
	replacePoiLinkAssociations(retrievedPointOfInterest, &pointOfInterestRelatedEntries.links, tx)

	replacePoiDetailAssociations(retrievedPointOfInterest, &pointOfInterestRelatedEntries.details, tx)

	replacePoiImageAssociations(retrievedPointOfInterest, &pointOfInterestRelatedEntries.images, tx)

	updatedPointOfInterest, err := repositories.UpdateEntity(retrievedPointOfInterest, tx)
	if err != nil {
		return nil, err
	}
	updatedPointOfInterest, err = repositories.UpdateColumn(updatedPointOfInterest, "support", putPointOfInterestBody.Support, tx)
	if err != nil {
		return nil, err
	}

	return updatedPointOfInterest, err
}

type pointOfInterestRelatedEntries struct {
	images     []*models.Image
	links      []*models.Link
	details    []*models.Detail
	categories []*models.Category
}

func CreateOrUpdatePointOfInterestRelatedEntities(postPointOfInterestBody *models.PostPointOfInterestBody, tx *gorm.DB) (*pointOfInterestRelatedEntries, error) {
	images, err := repositories.GetEntitiesByIDs[models.Image](&postPointOfInterestBody.ImageIDs, tx)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("error retrieving images")
	}
	if err == gorm.ErrRecordNotFound {
		images = &[]models.Image{}
	}

	// find or create links
	foundOrCreatedLinks := []*models.Link{}
	for _, postLinkBody := range postPointOfInterestBody.Links {

		foundOrCreatedLink, err := FirstOrCreateLink(&postLinkBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedLinks = append(foundOrCreatedLinks, foundOrCreatedLink)
	}

	// find or create details
	foundOrCreatedDetails := []*models.Detail{}
	for _, postDetailBody := range postPointOfInterestBody.Details {

		foundOrCreatedDetail, err := FirstOrCreateDetail(&postDetailBody, tx)
		if err != nil {
			return nil, err
		}
		foundOrCreatedDetails = append(foundOrCreatedDetails, foundOrCreatedDetail)
	}

	// find or create categories
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
		img := image
		imagesPointers = append(imagesPointers, &img)
	}

	return &pointOfInterestRelatedEntries{
		images:     imagesPointers,
		categories: foundOrCreatedCategories,
		details:    foundOrCreatedDetails,
		links:      foundOrCreatedLinks,
	}, nil
}

func CreatePointOfInterest(postPointOfInterestBody *models.PostPointOfInterestBody, authenticatedUser *models.User, tx *gorm.DB) (*models.PointOfInterest, error) {
	pointOfInterestRelatedEntries, err := CreateOrUpdatePointOfInterestRelatedEntities(postPointOfInterestBody, tx)
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
	updatedPointOfInterest, err := repositories.UpdateColumn(createdPointOfInterest, "support", postPointOfInterestBody.Support, tx)
	if err != nil {
		return nil, err
	}

	return updatedPointOfInterest, err
}

func replacePoiImageAssociations(poi *models.PointOfInterest, images *[]*models.Image, tx *gorm.DB) error {
	var oldImages []models.Image
	for _, image := range poi.Images {
		oldImages = append(oldImages, *image)
	}
	tx.Model(&poi).Association("Images").Delete(poi.Images)
	tx.Model(&poi).Association("Images").Append(images)
	for _, image := range oldImages {
		var count int64
		// TODO raw sql might just be better
		tx.Raw("SELECT COUNT(*) FROM points_of_interest_images WHERE image_id = ?", image.ID).Scan(&count)
		if count == 0 {
			tx.Delete(image, image)
		}

	}
	poi.Images = *images
	return nil
}

func replacePoiLinkAssociations(poi *models.PointOfInterest, links *[]*models.Link, tx *gorm.DB) error {
	var oldLinks []models.Link
	for _, link := range poi.Links {
		oldLinks = append(oldLinks, *link)
	}
	tx.Model(&poi).Association("Links").Delete(poi.Images)
	tx.Model(&poi).Association("Links").Append(links)
	for _, link := range oldLinks {
		var count int64
		// TODO raw sql might just be better
		tx.Raw("SELECT COUNT(*) FROM points_of_interest_links WHERE link_id = ?", link.ID).Scan(&count)
		if count == 0 {
			tx.Delete(link, link)
		}

	}
	poi.Links = *links
	return nil
}

func replacePoiDetailAssociations(poi *models.PointOfInterest, details *[]*models.Detail, tx *gorm.DB) error {
	var oldDetails []models.Detail
	for _, detail := range poi.Details {
		oldDetails = append(oldDetails, *detail)
	}
	tx.Model(&poi).Association("Details").Delete(poi.Details)
	tx.Model(&poi).Association("Details").Append(details)
	for _, detail := range oldDetails {
		var count int64
		// TODO raw sql might just be better
		tx.Raw("SELECT COUNT(*) FROM points_of_interest_details WHERE detail_id = ?", detail.ID).Scan(&count)
		if count == 0 {
			tx.Delete(detail, detail)
		}

	}
	poi.Details = *details
	return nil
}

func FindOrCreateOrUpdatePointOfInterest(postPointOfInterestBody *models.PostPointOfInterestBody, authenticatedUser *models.User, tx *gorm.DB) (*models.PointOfInterest, error) {
	// find
	getPointOfInterestBody := postPointOfInterestBody.ToGetPointOfInterestBody()

	pointOfInterest := getPointOfInterestBody.ToPointOfInterest()

	retrievedPointOfInterest, err := repositories.GetEntityByConditions(&pointOfInterest, tx)

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		// create

		createdPointOfInterest, err := CreatePointOfInterest(postPointOfInterestBody, authenticatedUser, tx)
		if err != nil {
			return nil, err
		}

		return createdPointOfInterest, err
	} else if retrievedPointOfInterest.UserID == authenticatedUser.ID {
		// update

		updatedPointOfInterest, err := UpdatePointOfInterest(retrievedPointOfInterest.ID, postPointOfInterestBody, authenticatedUser, tx)
		if err != nil {
			return nil, err
		}

		return updatedPointOfInterest, err
	} else {
		return retrievedPointOfInterest, nil
	}
}

// type PointOfInterestWithPosition struct {
// 	models.PointOfInterest
// 	position uint
// }

// func OrderPointsOfInterestByRoutePosition(routeID uint, pointsOfInterest *[]models.PointOfInterest) {

// 	derefrencedPointsOfInterest := *pointsOfInterest
// 	pointsOfInterestWithPositions
// 	sort.Slice(pointsOfInterest, func(i, j int) bool {
// 		poiA := derefrencedPointsOfInterest[i]
// 		poiB := derefrencedPointsOfInterest[j]
// 		routePointOfInterest := models.RoutesPointsOfInterest{
// 			RouteID:           routeID,
// 			PointOfInterestID: poiA.ID,
// 		}
// 		repositories.GetEntityByConditions[models.RoutesPointsOfInterest]()

// 	})

// }

func GetPointsOfInterestByRouteIDOrderedByPositionWithAssociations(routeID uint) (*[]models.PointOfInterest, error) {
	return repositories.GetPointsOfInterestByRouteIDOrderedByPositionWithAssociations(routeID, database.Db)
}
