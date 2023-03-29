package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetUsers() []models.User {
	var users []models.User

	result := database.Db.Find(&users)

	if result.Error != nil {
		println(result.Error)
	}

	return users
}
