package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetRoutes() []models.Route {
	var routes []models.Route

	result := database.Db.Find(&routes)

	if result.Error != nil {
		println(result.Error)
	}

	return routes
}
