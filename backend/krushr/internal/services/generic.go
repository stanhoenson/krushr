package services

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

// Singular
func GetEntityByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint) (*T, error) {
	return repositories.GetEntityByID[T](ID, database.Db)
}

func GetEntityByIDWithAssociations[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint, associations string) (*T, error) {
	return repositories.GetEntityByIDWithAssociations[T](ID, associations, database.Db)
}

func CreateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T) (*T, error) {
	return repositories.CreateEntity(entity, database.Db)
}

func DeleteEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T) (*T, error) {
	return repositories.DeleteEntity(entity, database.Db)
}

func DeleteEntityByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint) (*T, error) {
	return repositories.DeleteEntityByID[T](ID, database.Db)
}

func UpdateEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entity *T) (*T, error) {
	return repositories.UpdateEntity(entity, database.Db)
}

// Plural
func GetEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest]() (*[]T, error) {
	return repositories.GetEntities[T](database.Db)
}

func GetEntitiesWithAssociations[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](associations string) (*[]T, error) {
	return repositories.GetEntitiesWithAssociations[T](associations, database.Db)
}

func GetEntitiesByIDs[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](IDs *[]uint) (*[]T, error) {
	return repositories.GetEntitiesByIDs[T](IDs, database.Db)
}

func GetEntitiesByIDsAndConditions[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](IDs *[]uint, conditions *T) (*[]T, error) {
	return repositories.GetEntitiesByIDsAndConditions[T](IDs, conditions, database.Db)
}

func GetEntitiesByConditions[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](conditions *T) (*[]T, error) {
	return repositories.GetEntitiesByConditions[T](conditions, database.Db)
}

func CreateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entities *[]T) (*[]T, error) {
	return repositories.CreateEntities(entities, database.Db)
}

func DeleteEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entities *[]T) (*[]T, error) {
	return repositories.DeleteEntities(entities, database.Db)
}

func UpdateEntities[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](entities *[]T) (*[]T, error) {
	return repositories.UpdateEntities(entities, database.Db)
}
