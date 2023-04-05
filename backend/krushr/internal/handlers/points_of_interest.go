package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func getPointsOfInterest(c *gin.Context) {

	pointsOfInterest, err := services.GetPointsOfInterest()
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

	var newPointOfInterest models.PointOfInterest

	if err := c.BindJSON(&newPointOfInterest); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostPointOfInterest(&newPointOfInterest); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPoi, err := services.CreateEntity(&newPointOfInterest)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating point of interest"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdPoi)

}

func PointsOfInterestRoutes(r *gin.Engine) {
	routes := r.Group("/points-of-interest")
	{
		routes.GET("", getPointsOfInterest)
		routes.GET("/:id", getPointOfInterestByID)
		routes.POST("", postPointOfInterest)
	}
}
