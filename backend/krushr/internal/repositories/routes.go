package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetRoutes() (*[]models.Route, error) {
	var routes []models.Route

	result := database.Db.Find(&routes)

	if result.Error != nil {
		return nil, result.Error
	}

	return &routes, nil
}

func GetRouteById(ID uint) (*models.Route, error) {

	var route models.Route

	result := database.Db.First(route)

	if result.Error != nil {
		return nil, result.Error
	}
	return &route, nil
}
