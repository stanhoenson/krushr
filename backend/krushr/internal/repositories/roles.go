package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetRoles() []models.Role {
	var roles []models.Role

	result := database.Db.Find(&roles)

	if result.Error != nil {
		println(result.Error)
	}

	return roles
}
