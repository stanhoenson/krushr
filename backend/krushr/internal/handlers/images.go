package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func getImageDataByID(c *gin.Context) {

	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	image, err := services.GetEntityByID[models.Image](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving image"})
		return

	}

	c.File(image.Path)
	c.Status(200)

}

func postImage(c *gin.Context) {
	// single file
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error parsing MultipartForm" + err.Error()})
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
	r.GET("/imagedata/:id", getImageDataByID)

	routes := r.Group("/images")
	{
		routes.GET("/:id", func(ctx *gin.Context) {
			GetByID[models.Image](ctx)
		})
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByID[models.Image]))
		routes.POST("", wrappers.RoleWrapper(constants.Roles, postImage))
	}
}
