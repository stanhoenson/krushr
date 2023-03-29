package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetPointsOfInterest() []models.PointOfInterest {
	var pointsOfInterest []models.PointOfInterest

	result := database.Db.Find(&pointsOfInterest)

	if result.Error != nil {
		println(result.Error)
	}

	return pointsOfInterest
}
