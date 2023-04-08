package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetPointsOfInterest() (*[]models.PointOfInterest, error) {
	var pointsOfInterest []models.PointOfInterest

	result := database.Db.Find(&pointsOfInterest)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pointsOfInterest, nil
}

func GetPointOfInterest(ID uint) (*models.PointOfInterest, error) {
	var pointOfInterest models.PointOfInterest

	result := database.Db.First(&pointOfInterest, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pointOfInterest, nil
}
