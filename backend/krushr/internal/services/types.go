package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetTypes() []models.Type {
	types := repositories.GetTypes()
	return types
}
