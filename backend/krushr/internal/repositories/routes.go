package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func DeleteRouteByIDAndUserID(ID uint, userID uint) (*models.Route, error) {
	var route models.Route

	result := database.Db.Where("id = ?", ID).Where("user_id = ?", userID).Delete(&route)

	if result.Error != nil {
		return nil, result.Error
	}

	return &route, nil
}
