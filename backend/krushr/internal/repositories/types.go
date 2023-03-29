package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetTypes() []models.Type {
	var types []models.Type

	result := database.Db.Find(&types)

	if result.Error != nil {
		println(result.Error)
	}

	return types
}
