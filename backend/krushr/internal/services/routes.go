package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetRoutes() (*[]models.Route, error) {
	return repositories.GetRoutes()
}

func GetRouteByID(ID uint) (*models.Route, error) {
	return repositories.GetRouteById(ID)
}
