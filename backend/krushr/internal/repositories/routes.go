package repositories

import (
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/gorm"
)

func DeleteRouteByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route

	result := tx.Where("id = ?", ID).Where("user_id = ?", userID).Delete(&route)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil
}

func GetRouteByIDAndUserID(ID uint, userID uint, tx *gorm.DB) (*models.Route, error) {
	var route models.Route
	result := tx.Where("id = ?", ID).Where("user_id = ?", userID).First(&route)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil

}
