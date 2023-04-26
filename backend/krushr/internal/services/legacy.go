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
	legacyRoute, err := route.ToLegacyRoute()

	return legacyRoute, nil

}

func GetLegacyRoutes() (*[]models.LegacyRoute, error) {

	routes, err := GetEntitiesWithAssociations[models.Route](clause.Associations)
	if err != nil {
		return nil, err
	}

	var legacyRoutes []models.LegacyRoute
	for _, v := range *routes {

		fmt.Printf("%+v", v)
		legacyRoute, err := v.ToLegacyRoute()
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
