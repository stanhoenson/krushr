package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
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
	u64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		//return error
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	pointsOfInterest, err := services.GetPointOfInterest(ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving point of interest"})
		return
	}
	c.IndentedJSON(http.StatusOK, pointsOfInterest)
}

func postPointOfInterest(c *gin.Context) {

	var newPointOfInterest models.PointOfInterest

	if err := c.BindJSON(&newPointOfInterest); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating point of interest"})
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
