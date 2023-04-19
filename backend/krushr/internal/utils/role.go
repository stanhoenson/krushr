package utils

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) bool {
	value, exists := c.Get("authenticatedUser")
	user, ok := value.(*models.User)
	// TODO maybe get admin role and compare or something with ids
	if !ok || !exists || user.Role.Name != "Admin" {
		return false
	}
	return true
}

func IsAuthenticated(c *gin.Context) bool {
	return HasRole(c, constants.Roles)
}

func HasRole(c *gin.Context, roles []string) bool {
	value, exists := c.Get("authenticatedUser")
	user, ok := value.(*models.User)
	// TODO maybe get admin role and compare or something with ids
	if !ok || !exists || !StringArrayIncludesSubstring(roles, user.Role.Name) {
		return false
	}
	return true
}

func GetUserFromContext(c *gin.Context) (*models.User, error) {
	value, exists := c.Get("authenticatedUser")
	user, ok := value.(*models.User)
	if !exists || !ok {
		return nil, errors.New("No user in context")
	}
	return user, nil
}
