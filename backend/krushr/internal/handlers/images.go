package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	c.Header("Cache-Control", "no-store")

	c.File(image.Path)
	c.Status(200)
}

func postImage(c *gin.Context) {
	// single file
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error parsing MultipartForm"})
		return
	}

	createdFile, err := wrappers.WithTransaction(database.Db, func(tx *gorm.DB) (*models.Image, error) {
		return services.CreateImage(fileHeader, tx)
	})
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
		routes.GET("/:id", GetByIDDefault[models.Image])
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			DeleteByID(ctx, func(c *gin.Context, ID uint) (uint, error) {
				return wrappers.WithTransaction(database.Db, func(tx *gorm.DB) (uint, error) {
					return services.DeleteImage(ID, tx)
				})
			})
		}))
		routes.POST("", wrappers.RoleWrapper(constants.Roles, postImage))
	}
}
