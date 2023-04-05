package handlers

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func getCategories(c *gin.Context) {

	categories, err := services.GetEntites[models.Category]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving categories"})
		return
	}
	c.IndentedJSON(http.StatusOK, categories)
}

func postCategory(c *gin.Context) {
	var newCategory models.Category

	if err := c.BindJSON(&newCategory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostCategory(&newCategory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := services.CreateEntity(&newCategory)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating category"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdCategory)
}
func CategoriesRoutes(r *gin.Engine) {
	routes := r.Group("/categories")
	{
		routes.GET("", getCategories)
		routes.POST("", postCategory)
	}
}
