package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

func RegisterPointOfInterestRoutes(r *gin.Engine) {
	routes := r.Group("/points-of-interest")
	{
		routes.GET("", GetAll[models.PointOfInterest])
		routes.GET("/:id", GetByID[models.PointOfInterest])
		routes.DELETE("/:id", wrappers.RoleWrapper(constants.Roles, deletePointOfInterestByID))
		routes.PUT("/:id", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Put(ctx, func(requestBody *models.PutPointOfInterestBody) error {
				return validators.ValidatePutPointOfInterest(requestBody)
			}, func(c *gin.Context, ID uint, requestBody *models.PutPointOfInterestBody) (*models.PointOfInterest, error) {

				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return nil, err
				}

				return wrappers.WithTransaction(database.Db, func(tx *gorm.DB) (*models.PointOfInterest, error) {
					return services.UpdatePointOfInterest(ID, requestBody, user, tx)
				})
			})
		}))
		routes.POST("", wrappers.RoleWrapper(constants.Roles, func(c *gin.Context) {
			Post(c, func(requestBody *models.PostPointOfInterestBody) error {
				return validators.ValidatePostPointOfInterest(requestBody)
			}, func(c *gin.Context, requestBody *models.PostPointOfInterestBody) (*models.PointOfInterest, error) {
				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return nil, err
				}
				return wrappers.WithTransaction(database.Db, func(tx *gorm.DB) (*models.PointOfInterest, error) {
					return services.CreatePointOfInterest(requestBody, user, tx)
				})
			})
		}))
	}
}
