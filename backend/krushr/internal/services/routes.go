package services

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func DeleteRouteByIDAndAuthenticatedUser(ID uint, authenticatedUser *models.User) (*models.Route, error) {
	if authenticatedUser.Role.Role == constants.AdminRoleName {
		return repositories.DeleteEntityByID[models.Route](ID)
	}
	return repositories.DeleteRouteByIDAndUserID(ID, authenticatedUser.ID)
}

func CreateRoute(postRouteBody *models.PostRouteBody, authenticatedUser *models.User) (*models.Route, error) {
	route := models.Route{
		Title:    postRouteBody.Title,
		StatusID: postRouteBody.StatusID,
	}
	return repositories.CreateEntity(&route)
}
