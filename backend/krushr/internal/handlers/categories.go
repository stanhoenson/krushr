package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func putCategory(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)
	var putCategoryBody models.PutCategoryBody

	if err := c.BindJSON(&putCategoryBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePutCategoryBody(&putCategoryBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = services.UpdateCategory(ID, &putCategoryBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating category"})
		return
	}

	c.IndentedJSON(http.StatusOK, putCategoryBody)
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
	categories, err := services.GetEntities[models.Category]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving categories"})
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
}

func postCategory(c *gin.Context) {
	var postCategoryBody models.PostCategoryBody

	if err := c.BindJSON(&postCategoryBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostCategoryBody(&postCategoryBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := services.CreateCategory(&postCategoryBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating category"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdCategory)
}

func RegisterCategoryRoutes(r *gin.Engine) {
	routes := r.Group("/categories")
	{
		routes.GET("", getCategories)
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, deleteCategoryByID))
		routes.POST("", wrappers.RoleWrapper(constants.Roles, postCategory))
		routes.PUT("", wrappers.RoleWrapper(constants.Roles, putCategory))
	}
}
