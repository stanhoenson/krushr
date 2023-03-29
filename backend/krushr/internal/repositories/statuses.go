package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetStatuses() []models.Status {
	var statuses []models.Status

	result := database.Db.Find(&statuses)

	if result.Error != nil {
		println(result.Error)
	}

	return statuses
}
