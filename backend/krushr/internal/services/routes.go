package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetRoutes() []models.Route {

	var routes = repositories.GetRoutes()
	return routes
}
