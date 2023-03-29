package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetCategories() []models.Category {
	var categories []models.Category

	result := database.Db.Find(&categories)

	if result.Error != nil {
		println(result.Error)
	}

	return categories
}
