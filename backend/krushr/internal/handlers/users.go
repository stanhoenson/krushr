package handlers

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	routes := r.Group("/users")
	{
		routes.GET("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, GetAllDefault[models.User]))
		routes.POST("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			Post(ctx, func(c *gin.Context, requestBody *models.PostUserBody) error {
				return nil
			}, func(c *gin.Context, requestBody *models.PostUserBody) (*models.User, error) {
				return services.CreateUser(requestBody, database.Db)
			})
		}))
		routes.PUT("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			Put(ctx, func(c *gin.Context, requestBody *models.PutUserBody) error {
				return nil
			}, func(c *gin.Context, ID uint, requestBody *models.PutUserBody) (*models.User, error) {
				return services.UpdateUser(ID, requestBody, database.Db)
			})
		}))
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByIDDefault[models.User]))
	}
}
