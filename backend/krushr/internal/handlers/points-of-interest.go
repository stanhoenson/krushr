package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func putPointOfInterest(c *gin.Context) {
	var putPointOfInterestBody models.PutPointOfInterestBody

	if err := c.BindJSON(&putPointOfInterestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePutPointOfInterest(&putPointOfInterestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user in context"})
		return
	}

	updatedPointOfInterest, err := services.UpdatePointOfInterest(&putPointOfInterestBody, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating point of interest"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedPointOfInterest)
}

func deletePointOfInterestByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user in context"})
		return
	}

	deletedPointOfInterest, err := services.DeletePointOfInterestByIDAndAuthentictedUser(ID, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting point of interest"})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedPointOfInterest)
}

func getPointsOfInterest(c *gin.Context) {
	pointsOfInterest, err := services.GetEntities[models.PointOfInterest]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving points of interest"})
		return
	}

	c.IndentedJSON(http.StatusOK, pointsOfInterest)
}

func getPointOfInterestByID(c *gin.Context) {
	id := c.Param("id")

	// Convert string to uint
	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	pointsOfInterest, err := services.GetEntity[models.PointOfInterest](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving point of interest"})
		return
	}

	c.IndentedJSON(http.StatusOK, pointsOfInterest)
}

func postPointOfInterest(c *gin.Context) {
	var postPointOfInterestBody models.PostPointOfInterestBody

	if err := c.BindJSON(&postPointOfInterestBody); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostPointOfInterest(&postPointOfInterestBody); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user in context"})
		return
	}

	createdPointOfInterest, err := services.CreatePointOfInterest(&postPointOfInterestBody, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating point of interest"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdPointOfInterest)
}

func RegisterPointOfInterestRoutes(r *gin.Engine) {
	routes := r.Group("/points-of-interest")
	{
		routes.GET("", getPointsOfInterest)
		routes.GET("/:id", getPointOfInterestByID)
		routes.DELETE("/:id", wrappers.RoleWrapper(constants.Roles, deletePointOfInterestByID))
		routes.PUT("", wrappers.RoleWrapper(constants.Roles, putPointOfInterest))
		routes.POST("", wrappers.RoleWrapper(constants.Roles, postPointOfInterest))
	}
}
