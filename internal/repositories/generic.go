package repositories

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"

	"gorm.io/gorm"
)

// Singular
func GetEntityByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint, tx *gorm.DB) (*T, error) {
	var entity T

	result := tx.First(&entity, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func GetEntityByIDWithAssociations[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint, associations []string, tx *gorm.DB) (*T, error) {
	var entity T

	var transaction *gorm.DB = tx

	for _, association := range associations {
		transaction = transaction.Preload(association)
	}
	result := transaction.First(&entity, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &entity, nil
}

func GetEntityByConditions[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T, tx *gorm.DB) (*T, error) {
	result := tx.First(&entity, &entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
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

	if result.RowsAffected == 0 {
		return nil, errors.New("no rows affected by delete operation")
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

func DeleteEntityByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint, tx *gorm.DB) (uint, error) {
	var entity T

	result := tx.Delete(&entity, ID)

	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected == 0 {
		return 0, errors.New("no rows affected by delete operation")
	}

	return ID, nil
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

func GetEntitiesWithAssociations[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](associations []string, tx *gorm.DB) (*[]T, error) {
	var entities []T

	var transaction *gorm.DB = tx

	for _, association := range associations {
		transaction = transaction.Preload(association)
	}
	result := transaction.Find(&entities)

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
	} else if len(entities) != len(*IDs) {
		return nil, gorm.ErrRecordNotFound
	}

	return &entities, nil
}

func GetEntitiesByIDsAndConditions[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](IDs *[]uint, conditions *T, tx *gorm.DB) (*[]T, error) {
	var entities []T

	result := tx.Find(&entities, IDs, conditions)

	if result.Error != nil {
		return nil, result.Error
	} else if len(entities) != len(*IDs) {
		return nil, gorm.ErrRecordNotFound
	}

	return &entities, nil
}

func GetEntitiesByConditions[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](conditions *T, tx *gorm.DB) (*[]T, error) {
	var entities []T

	result := tx.Find(&entities, conditions)

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

	if result.RowsAffected == 0 {
		return nil, errors.New("no rows affected by delete operation")
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

func UpdateColumn[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T, column string, value any, tx *gorm.DB) (*T, error) {
	result := tx.Model(&entity).Update(column, value)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}
