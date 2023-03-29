package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetStatuses() []models.Status {
	statuses := repositories.GetStatuses()
	return statuses
}
