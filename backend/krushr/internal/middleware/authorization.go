package middleware

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt := c.GetHeader("Authorization")
		if jwt == "" {
			c.Set("authenticatedUser", nil)
		} else {
			user, err := services.GetUserFromJWT(jwt)
			if err != nil {
				// TODO maybe StatusBadRequest
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			c.Set("authenticatedUser", user)
		}
		c.Next()
	}
}
