package repositories

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
)

// Singular
func GetEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](ID uint) (*T, error) {
	var entity T

	result := database.Db.First(&entity, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func CreateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entity *T) (*T, error) {
	result := database.Db.Create(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func DeleteEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entity *T) (*T, error) {
	result := database.Db.Delete(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func DeleteEntityByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](ID uint) (*T, error) {
	var entity T

	result := database.Db.Delete(&entity, ID)
	fmt.Println(result.RowsAffected)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func UpdateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entity *T) (*T, error) {
	result := database.Db.Updates(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

// Plural
func GetEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role]() (*[]T, error) {
	var entities []T

	result := database.Db.Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entities, nil
}

func CreateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entities *[]T) (*[]T, error) {
	result := database.Db.Create(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func DeleteEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entities *[]T) (*[]T, error) {
	result := database.Db.Delete(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func UpdateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role](entities *[]T) (*[]T, error) {
	result := database.Db.Updates(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}
