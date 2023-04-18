package wrappers

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/gin-gonic/gin"
)

func RoleWrapper(roles []string, handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		hasRole := utils.HasRole(c, roles)
		if hasRole {
			handler(c)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
