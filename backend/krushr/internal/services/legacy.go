package services

import (
	"fmt"

	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/gorm/clause"
)

func GetLegacyRouteByID(ID uint) (*models.LegacyRoute, error) {
	route, err := GetEntityByIDWithAssociations[models.Route](ID, clause.Associations)
	if err != nil {
		return nil, err
	}
	pointsOfInterest, err := GetPointsOfInterestByRouteIDOrderedByPositionWithAssociations(route.ID)
	if err != nil {
		return nil, err
	}

	if len(route.PointsOfInterest) != len(*pointsOfInterest) {
		return nil, fmt.Errorf("error retrieving legacyRoute")
	}

	for index := range route.PointsOfInterest {
		route.PointsOfInterest[index] = &(*pointsOfInterest)[index]
	}

	legacyRoute, err := route.ToLegacyRoute(true)
	if err != nil {
		return nil, err
	}

	return legacyRoute, nil
}

func GetLegacyRoutes() (*[]models.LegacyRoute, error) {
	routes, err := GetEntitiesWithAssociations[models.Route](clause.Associations)
	if err != nil {
		return nil, err
	}

	var legacyRoutes []models.LegacyRoute
	for _, v := range *routes {

		legacyRoute, err := v.ToLegacyRoute(false)
		if err != nil {
			return nil, err
		}
		legacyRoutes = append(legacyRoutes, *legacyRoute)
	}

	return &legacyRoutes, nil
}

func GetLegacyMenus() (*[]models.LegacyMenu, error) {
	categories, err := GetEntities[models.Category]()
	if err != nil {
		return nil, err
	}
	var legacyMenus []models.LegacyMenu
	for _, v := range *categories {
		legacyMenus = append(legacyMenus, v.ToLegacyMenu())
	}
	return &legacyMenus, nil
}
