package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetEntries() []models.Entry {
	var entries []models.Entry

	result := database.Db.Find(&entries)

	if result.Error != nil {
		println(result.Error)
	}

	return entries
}
