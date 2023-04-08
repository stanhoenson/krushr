package handlers

import (
	"fmt"
	"net/http"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(c)
	fmt.Println(newUser)

	if err := validators.ValidateSignUp(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateUser(&newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error signing up"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdUser)
}

func AuthenticationRoutes(r *gin.Engine) {
	routes := r.Group("/authentication")
	{
		// routes.POST("/sign-in", signIn)
		routes.POST("/sign-up", signUp)
	}
}
