package handlers

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func postImage(c *gin.Context) {
	// single file
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error parsing MultipartForm"})
		return
	}

	createdFile, err := services.CreateImage(fileHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating image"})
		return

	}

	c.IndentedJSON(http.StatusOK, createdFile)
}

func RegisterImageRoutes(r *gin.Engine) {
	routes := r.Group("/images")
	{
		routes.GET("/:id", getPointOfInterestByID)
		routes.DELETE("/:id", deletePointOfInterestByID)
		routes.POST("", postPointOfInterest)
	}
}
