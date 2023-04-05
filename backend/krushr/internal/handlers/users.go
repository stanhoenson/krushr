package handlers

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {

	users, err := services.GetEntites[models.User]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving users"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func postUser(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostUser(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := services.CreateEntity(&newUser)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdUser)
}

func UsersRoutes(r *gin.Engine) {
	routes := r.Group("/users")
	{
		routes.GET("", getUsers)
		routes.POST("", postUser)
	}
}
