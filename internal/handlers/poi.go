package handlers

import (
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

func RegisterPointOfInterestRoutes(r *gin.Engine) {
	routes := r.Group("/points-of-interest")
	{
		routes.GET("", GetAllDefault[models.PointOfInterest])
		routes.GET("/:id", GetByIDDefault[models.PointOfInterest])
		routes.DELETE("/:id", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			DeleteByID(ctx, func(c *gin.Context, ID uint) (uint, error) {
				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return 0, err
				}

				return services.DeletePointOfInterestByIDAndAuthentictedUser(ID, user)
			})
		}))
		routes.PUT("/:id", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Put(ctx, func(c *gin.Context, requestBody *models.PutPointOfInterestBody) error {
				return validators.ValidatePostPointOfInterestBody(requestBody)
			},
				func(c *gin.Context, ID uint, requestBody *models.PutPointOfInterestBody) (*models.PointOfInterest, error) {
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
			Post(c, func(c *gin.Context, requestBody *models.PostPointOfInterestBody) error {
				return validators.ValidatePostPointOfInterestBody(requestBody)
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
