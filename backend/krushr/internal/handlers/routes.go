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

func DeleteRouteByID(c *gin.Context) {
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

	deletedRoute, err := services.DeleteRouteByIDAndAuthenticatedUser(ID, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting route"})
		return
	}

	c.JSON(http.StatusOK, deletedRoute)
}

func RegisterRouteRoutes(r *gin.Engine) {
	routes := r.Group("/routes")
	{
		routes.GET("", GetAll[models.Route])
		routes.GET("/:id", GetByID[models.Route])
		routes.POST("", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Post(ctx, func(requestBody *models.PostRouteBody) error {
				return validators.ValidatePostRouteBody(requestBody)
			}, func(c *gin.Context, requestBody *models.PostRouteBody) (*models.Route, error) {
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
			Put(ctx, func(requestBody *models.PutRouteBody) error {
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
		routes.DELETE("/:id", wrappers.RoleWrapper(constants.Roles, DeleteRouteByID))
	}
}
