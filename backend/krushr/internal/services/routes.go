package services

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
	"github.com/stanhoenson/krushr/internal/utils"
)

func DeleteRouteByIDAndAuthenticatedUser(ID uint, authenticatedUser *models.User) (*models.Route, error) {
	if authenticatedUser.Role.Name == constants.AdminRoleName {
		return repositories.DeleteEntityByID[models.Route](ID)
	}
	return repositories.DeleteRouteByIDAndUserID(ID, authenticatedUser.ID)
}

func CreateRoute(postRouteBody *models.PostRouteBody, authenticatedUser *models.User) (*models.Route, error) {


    pointsOfInterest := GetEntitiesByIDs[models.PointOfInterest](post)

	route := models.Route{
		Name:     postRouteBody.Name,
		StatusID: postRouteBody.StatusID,
	}
	return repositories.CreateEntity(&route)
}
