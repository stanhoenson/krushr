package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func GetPointsOfInterest(c *gin.Context) {

	pointsOfInterest, err := services.GetPointsOfInterest()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving points of interest"})
		return
	}
	c.IndentedJSON(http.StatusOK, pointsOfInterest)
}

func GetPointOfInterestByID(c *gin.Context) {
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

func PointsOfInterestRoutes(r *gin.Engine) {
	routes := r.Group("/points-of-interest")
	{
		routes.GET("", GetPointsOfInterest)
		routes.GET("/:id", GetPointOfInterestByID)
	}
}
