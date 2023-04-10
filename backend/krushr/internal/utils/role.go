package utils

import (
	"errors"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) error {

	entry, exists := c.Get("authenticatedUser")
	user, ok := entry.(*models.User)
	//TODO maybe get admin role and compare or something with ids
	if !ok || !exists || user.Role.Role != "Admin" {
		return errors.New("No admin")
	}
	return nil

}
