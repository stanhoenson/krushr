package handlers

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
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

	c.IndentedJSON(http.StatusOK, jwt)
}

func RegisterAuthenticationRoutes(r *gin.Engine) {
	routes := r.Group("/authentication")
	{
		routes.POST("/sign-in", signIn)
		routes.POST("/sign-up", signUp)
	}
}
