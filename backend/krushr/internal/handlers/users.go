package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func putUser(c *gin.Context) {
	var updatedUser models.User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePutUser(&updatedUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.UpdateEntity(&updatedUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating user"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

func deleteUserByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedUser, err := services.DeleteEntityByID[models.User](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting user"})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedUser)
}

func getUsers(c *gin.Context) {
	users, err := services.GetEntities[models.User]()
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

func RegisterUserRoutes(r *gin.Engine) {
	routes := r.Group("/users")
	{
		routes.GET("", getUsers)
		routes.POST("", postUser)
		routes.PUT("", putUser)
		routes.DELETE("/:id", deleteUserByID)
	}
}
