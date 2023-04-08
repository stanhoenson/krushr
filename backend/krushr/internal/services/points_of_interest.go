package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetPointsOfInterest() (*[]models.PointOfInterest, error) {
	return repositories.GetPointsOfInterest()
}

func GetPointOfInterest(ID uint) (*models.PointOfInterest, error) {
	return repositories.GetPointOfInterest(ID)
}
