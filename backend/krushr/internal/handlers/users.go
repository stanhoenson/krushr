package handlers

import (
	"fmt"
	"net/http"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetMe(c *gin.Context) {
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving " + utils.GetTypeString(&models.User{})})
	}
	entity, err := services.GetEntityByIDWithAssociations[models.User](user.ID, clause.Associations)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving " + utils.GetTypeString(entity)})
		return
	}

	c.IndentedJSON(http.StatusOK, entity)
}

func RegisterUserRoutes(r *gin.Engine) {
	routes := r.Group("/users")
	{
		routes.GET("/me", wrappers.RoleWrapper(constants.Roles, GetMe))
		routes.GET("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, GetAllDefault[models.User]))
		routes.POST("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			Post(ctx, func(c *gin.Context, requestBody *models.PostUserBody) error {
				return nil
			}, func(c *gin.Context, requestBody *models.PostUserBody) (*models.User, error) {
				return services.CreateUser(requestBody, database.Db)
			})
		}))
        routes.PUT("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			Put(ctx, func(c *gin.Context, requestBody *models.PutUserBody) error {
				return nil
			}, func(c *gin.Context, ID uint, requestBody *models.PutUserBody) (*models.User, error) {
				return services.UpdateUser(ID, requestBody, database.Db)
			})
		}))
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			DeleteByID(ctx, func(c *gin.Context, ID uint) (*models.User, error) {
				user, err := utils.GetUserFromContext(c)
				if err != nil {
					return nil, err
				}

				if ID == user.ID {
					return nil, fmt.Errorf("you cannot delete yourself")
				}

				return services.DeleteEntityByID[models.User](ID)
			})
		}))
	}
}
