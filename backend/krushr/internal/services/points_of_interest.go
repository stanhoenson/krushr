package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetPointsOfInterest() []models.PointOfInterest {
	pointsOfInterest := repositories.GetPointsOfInterest()
	return pointsOfInterest
}
