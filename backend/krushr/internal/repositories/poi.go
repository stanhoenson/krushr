package repositories

import (
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DeletePointOfInterestByIDAndAuthentictedUser(ID uint, userID uint, tx *gorm.DB) (*models.PointOfInterest, error) {
	var pointOfInterest models.PointOfInterest

	result := tx.Where("id = ?", ID).Where("user_id = ?", userID).Delete(&pointOfInterest)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pointOfInterest, nil
}

func GetPointOfInterestByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (*models.PointOfInterest, error) {
	var pointOfInterest models.PointOfInterest
	result := tx.Where("id = ?", ID).Where("user_id = ?", userID).First(&pointOfInterest)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pointOfInterest, nil
}

func GetPointsOfInterestByRouteIDOrderedByPositionWithAssociations(routeID uint, tx *gorm.DB) (*[]models.PointOfInterest, error) {
	var pointsOfInterest []models.PointOfInterest

	result := tx.Preload(clause.Associations).Joins("JOIN routes_points_of_interest ON routes_points_of_interest.point_of_interest_id = points_of_interest.id").
		Where("routes_points_of_interest.route_id = ?", routeID).
		Order("routes_points_of_interest.position ASC").
		Find(&pointsOfInterest)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pointsOfInterest, nil
}
