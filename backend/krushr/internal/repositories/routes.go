package repositories

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DeleteRouteByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route

	result := tx.Where("id = ?", ID).Where("user_id = ?", userID).Delete(&route)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil
}

func GetRoutesWithAssociationsByUserID(userID uint, tx *gorm.DB) (*[]models.Route, error) {
	var routes []models.Route

	result := tx.Preload(clause.Associations).Joins("Status").Where("user_id = ? OR (user_id != ? AND status.name = ? )", userID, userID, constants.PublishedStatusName).Find(&routes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &routes, nil

}

func GetPublishedRoutes(tx *gorm.DB) (*[]models.Route, error) {
	var routes []models.Route

	result := tx.Preload(clause.Associations).Joins("Status").Where(" status.name = ? ", constants.PublishedStatusName).Find(&routes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &routes, nil

}

func GetRouteByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route
	result := tx.Where("id = ?", ID).Where("user_id = ?", userID).First(&route)

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
