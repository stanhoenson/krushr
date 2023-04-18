package repositories

import (
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/gorm"
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
