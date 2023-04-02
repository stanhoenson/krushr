package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetCategories() []models.Category {
	categories := repositories.GetCategories()
	return categories
}

