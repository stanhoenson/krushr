package repositories

import (
	"github.com/stanhoenson/krushr/internal/models"

	"gorm.io/gorm"
)

// Singular
func GetEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint, tx *gorm.DB) (*T, error) {
	var entity T

	result := tx.First(&entity, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func CreateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T, tx *gorm.DB) (*T, error) {
	result := tx.Create(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func DeleteEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T, tx *gorm.DB) (*T, error) {
	result := tx.Delete(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func FirstOrCreateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T, cond *T, tx *gorm.DB) (*T, error) {
	result := tx.FirstOrCreate(&entity, cond)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

func DeleteEntityByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint, tx *gorm.DB) (*T, error) {
	var entity T

	result := tx.Delete(&entity, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func UpdateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T, tx *gorm.DB) (*T, error) {
	result := tx.Updates(&entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}

// Plural
func GetEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](tx *gorm.DB) (*[]T, error) {
	var entities []T

	result := tx.Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entities, nil
}

func GetEntitiesByIDs[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](IDs *[]uint, tx *gorm.DB) (*[]T, error) {
	var entities []T

	result := tx.Find(&entities, IDs)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entities, nil
}

func CreateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entities *[]T, tx *gorm.DB) (*[]T, error) {
	result := tx.Create(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func DeleteEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entities *[]T, tx *gorm.DB) (*[]T, error) {
	result := tx.Delete(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func UpdateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entities *[]T, tx *gorm.DB) (*[]T, error) {
	result := tx.Updates(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}
