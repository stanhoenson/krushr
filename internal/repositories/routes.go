package repositories

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DeleteRouteByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (uint, error) {
	var route models.Route

	result := tx.Where("id = ?", ID).Where("user_id = ?", userID).Delete(&route)

	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("no route found for the given ID and userID")
	}

	return ID, nil
}

func GetRoutesWithAssociationsByUserID(userID uint, tx *gorm.DB) (*[]models.Route, error) {
	var routes []models.Route

	result := tx.Preload(clause.Associations).Joins("Status").Where("routes.user_id = ? OR (routes.user_id != ? AND status.name = ? )", userID, userID, constants.PublishedStatusName).Find(&routes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &routes, nil
}

func GetPublishedRoutes(tx *gorm.DB) (*[]models.Route, error) {
	var routes []models.Route

	result := tx.
		Preload("Images").
		Preload("Details").
		Preload("Links").
		Preload("Categories").
		Joins("LEFT JOIN statuses ON routes.status_id = statuses.id").
		Where("statuses.name = ?", constants.PublishedStatusName).
		Find(&routes)
	if result.Error != nil {
		return nil, result.Error
	}

	return &routes, nil
}

func GetPublishedRouteByID(ID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route

	result := tx.Preload(clause.Associations).Joins("Status").Where("status.name = ? AND routes.id = ?", constants.PublishedStatusName, ID).First(&route)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil
}

func GetPublishedRouteByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route

	result := tx.Preload(clause.Associations).Joins("Status").Where("status.name = ? AND routes.id = ? OR (routes.id = ? AND routes.user_id = ?)", constants.PublishedStatusName, ID, ID, userID).First(&route)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil
}

func GetRouteByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route
	result := tx.Preload(clause.Associations).Where("id = ?", ID).Where("user_id = ?", userID).First(&route)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil
}

func GetRouteByIDWithAssociations(ID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route
	result := tx.Preload(clause.Associations).First(&route, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil
}
