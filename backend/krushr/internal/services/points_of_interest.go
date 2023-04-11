package services

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/repositories"
)

func DeletePointOfInterestByID(authenticatedUser *models.User, ID uint) {
	if authenticatedUser.Role.Role == constants.AdminRoleName {
		repositories.DeleteEntityByID[models.PointOfInterest](ID)
	} else {
	}
}
