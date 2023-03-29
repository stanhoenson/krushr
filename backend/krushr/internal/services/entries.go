package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func GetEntries() []models.Entry {
	entries := repositories.GetEntries()
	return entries
}
