package middleware

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := c.Request.Cookie("jwt")
		if err != nil {
			// no cookie found
			c.Set("authenticatedUser", nil)
		} else if jwt.Value == "" || jwt.Valid() != nil {
			// invalid cookie set to null
			c.Set("authenticatedUser", nil)
			c.SetCookie("jwt", "", 0, "/", env.Domain, true, true)
		} else {
			user, err := services.GetUserFromJWT(jwt.Value)
			if err != nil {
				c.SetCookie("jwt", "", 0, "/", env.Domain, true, true)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			// gin.DefaultWriter.Write([]byte("authenticated"))
			c.Set("authenticatedUser", user)
		}
		c.Next()
	}
}
