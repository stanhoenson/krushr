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
	"gorm.io/gorm/clause"
)

func RegisterRouteRoutes(r *gin.Engine) {
	routes := r.Group("/routes")
	{
		routes.GET("", func(c *gin.Context) {
			GetAll(c, func(c *gin.Context) (*[]models.Route, error) {
				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return services.GetPublishedRoutes()
				}

				if utils.HasRole(c, []string{constants.AdminRoleName}) {
					return services.GetEntitiesWithAssociations[models.Route](clause.Associations)
				}

				return services.GetRoutesWithAssociationsByUserID(user.ID)
			})
		})
		routes.GET("/:id", func(ctx *gin.Context) {
			GetByID(ctx, func(c *gin.Context, ID uint) (*models.Route, error) {
				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return services.GetPublishedRouteByID(ID)
				}

				if utils.HasRole(c, []string{constants.AdminRoleName}) {
					return services.GetRouteByIDWithAssociations(ID)
				}

				return services.GetPublishedRouteByIDAndUserID(ID, user.ID)
			})
		})
		routes.POST("", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Post(ctx, func(c *gin.Context, requestBody *models.PostRouteBody) error {
				return validators.ValidatePostRouteBody(requestBody)
			},
				func(c *gin.Context, requestBody *models.PostRouteBody) (*models.Route, error) {
					user, err := utils.GetUserFromContext(c)
					if err != nil {
						return nil, err
					}
					return wrappers.WithTransaction(database.Db, func(tx *gorm.DB) (*models.Route, error) {
						return services.CreateRoute(requestBody, user, tx)
					})
				})
		}))
		routes.PUT("/:id", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Put(ctx, func(c *gin.Context, requestBody *models.PutRouteBody) error {
				return validators.ValidatePutRoute(requestBody)
			}, func(c *gin.Context, ID uint, requestBody *models.PutRouteBody) (*models.Route, error) {
				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return nil, err
				}
				return wrappers.WithTransaction(database.Db, func(tx *gorm.DB) (*models.Route, error) {
					return services.UpdateRoute(ID, requestBody, user, tx)
				})
			})
		}))
		routes.DELETE("/:id", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			DeleteByID(ctx, func(c *gin.Context, ID uint) (*models.Route, error) {
				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return nil, err
				}

				return services.DeleteRouteByIDAndAuthenticatedUser(ID, user)
			})
		}))
	}
}
