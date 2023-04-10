package middleware

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Set("authenticatedUser", nil)

		} else {
			user, err := services.GetUserFromJwt(authHeader)
			if err != nil {
				//TODO maybe StatusBadRequest
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized" + err.Error()})
				return
			}
			c.Set("authenticatedUser", user)
		}
		c.Next()

	}
}
