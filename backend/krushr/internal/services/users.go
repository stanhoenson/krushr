package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetUsers() []models.User {
	users := repositories.GetUsers()
	return users
}
