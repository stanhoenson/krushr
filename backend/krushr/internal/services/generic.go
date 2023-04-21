package services

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

// Singular
func GetEntity[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](ID uint) (*T, error) {
	return repositories.GetEntity[T](ID, database.Db)
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

func GetEntitiesByIDs[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](IDs *[]uint) (*[]T, error) {
	return repositories.GetEntitiesByIDs[T](IDs, database.Db)
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
