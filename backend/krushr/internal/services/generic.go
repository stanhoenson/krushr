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

	return repositories.CreateEntity[T](entity)
}

func DeleteEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {

	return repositories.DeleteEntity[T](entity)
}

func UpdateEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {

	return repositories.UpdateEntity[T](entity)
}

// Plural

func GetEntites[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User]() (*[]T, error) {

	return repositories.GetEntites[T]()
}

func CreateEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {

	return repositories.CreateEntities[T](entities)
}

func DeleteEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {
	return repositories.DeleteEntities[T](entities)
}

func UpdateEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {
	return repositories.UpdateEntities[T](entities)
}
