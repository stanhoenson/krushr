package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

// Singular
func GetEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](ID uint) (*T, error) {
	return repositories.GetEntity[T](ID)
}

func CreateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entity *T) (*T, error) {
	return repositories.CreateEntity(entity)
}

func DeleteEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entity *T) (*T, error) {
	return repositories.DeleteEntity(entity)
}

func DeleteEntityByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](ID uint) (*T, error) {
	return repositories.DeleteEntityByID[T](ID)
}

func UpdateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entity *T) (*T, error) {
	return repositories.UpdateEntity(entity)
}

// Plural
func GetEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role]() (*[]T, error) {
	return repositories.GetEntities[T]()
}
func GetEntitiesByIDs[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](IDs *[]uint) (*[]T, error) {
	return repositories.GetEntitiesByIDs[T](IDs)
}

func CreateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entities *[]T) (*[]T, error) {
	return repositories.CreateEntities(entities)
}

func DeleteEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entities *[]T) (*[]T, error) {
	return repositories.DeleteEntities(entities)
}

func UpdateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entities *[]T) (*[]T, error) {
	return repositories.UpdateEntities(entities)
}
