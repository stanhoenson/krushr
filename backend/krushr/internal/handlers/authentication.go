package handlers

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var signUpBody models.SignUpBody

	if err := c.BindJSON(&signUpBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateUserFromSignUpBody(&signUpBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error signing up"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdUser)
}

func signIn(c *gin.Context) {
	var signInBody models.SignInBody

	if err := c.BindJSON(&signInBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := services.Authenticate(&signInBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error signing in"})
		return
	}

	c.SetCookie("jwt", jwt, int(constants.TokenValidityPeriod.Seconds()), "/", env.Domain, true, true)

	c.Status(http.StatusOK)
}

func signOut(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "/", env.Domain, true, true)

	c.Status(http.StatusOK)
}

func RegisterAuthenticationRoutes(r *gin.Engine) {
	routes := r.Group("/authentication")
	{
		routes.POST("/sign-in", signIn)
		routes.GET("/sign-out", signOut)
		routes.POST("/sign-up", wrappers.RoleWrapper([]string{constants.AdminRoleName}, signUp))
	}
}
