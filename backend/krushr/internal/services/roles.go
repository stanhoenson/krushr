package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetRoles() []models.Role {
	roles := repositories.GetRoles()
	return roles
}
