package services

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

//Singular

func GetEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](ID uint) (*T, error) {

	return repositories.GetEntity[T](ID)
}

func CreateEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {

	return repositories.CreateEntity(entity)
}

func DeleteEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {

	return repositories.DeleteEntity(entity)
}

func DeleteEntityByID[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](ID uint) (*T, error) {

	return repositories.DeleteEntityByID[T](ID)
}

func UpdateEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {

	return repositories.UpdateEntity(entity)
}

// Plural

func GetEntites[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User]() (*[]T, error) {

	return repositories.GetEntities[T]()
}

func CreateEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {

	return repositories.CreateEntities(entities)
}

func DeleteEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {
	return repositories.DeleteEntities(entities)
}

func UpdateEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {
	return repositories.UpdateEntities(entities)
}
