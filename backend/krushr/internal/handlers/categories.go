package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func putCategory(c *gin.Context) {
	var updatedCategory models.Category

	if err := c.BindJSON(&updatedCategory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePutCategory(&updatedCategory); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.UpdateEntity(&updatedCategory)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating category"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedCategory)
}

func deleteCategoryByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedCategory, err := services.DeleteEntityByID[models.Category](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting category"})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedCategory)
}

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
		routes.DELETE("/:id", deleteCategoryByID)
		routes.POST("", postCategory)
		routes.PUT("", putCategory)
	}
}
