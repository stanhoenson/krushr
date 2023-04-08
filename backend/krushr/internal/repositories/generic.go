package repositories

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

// Singular
func GetEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](ID uint) (*T, error) {
	var entity T

	result := database.Db.First(&entity, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func CreateEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {
	result := database.Db.Create(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func DeleteEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {
	result := database.Db.Delete(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func DeleteEntityByID[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](ID uint) (*T, error) {
	var entity T

	result := database.Db.Delete(&entity, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func UpdateEntity[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entity *T) (*T, error) {
	result := database.Db.Updates(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

// Plural
func GetEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User]() (*[]T, error) {
	var entities []T

	result := database.Db.Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entities, nil
}

func CreateEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {
	result := database.Db.Create(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func DeleteEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {
	result := database.Db.Delete(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func UpdateEntities[T models.Category | models.Entry | models.PointOfInterest | models.Role | models.Route | models.Status | models.Type | models.User](entities *[]T) (*[]T, error) {
	result := database.Db.Updates(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}
