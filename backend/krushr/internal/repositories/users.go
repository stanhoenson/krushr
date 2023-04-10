package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := database.Db.Where("email = ?", email).First(&user)

	if result.Error != nil {

		return nil, result.Error
	}

	return &user, nil
}

func GetUserByIDWithRole(ID uint) (*models.User, error) {

	var user models.User

	result := database.Db.Preload("Role").First(&user, ID)

	if result.Error != nil {

		return nil, result.Error
	}

	return &user, nil
}
