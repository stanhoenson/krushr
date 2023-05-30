package handlers

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	routes := r.Group("/categories")
	{
		routes.GET("", GetAllDefault[models.Category])
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByIDDefault[models.Category]))
		routes.POST("", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Post(ctx, func(c *gin.Context, requestBody *models.PostCategoryBody) error {
				return validators.ValidatePostCategoryBody(requestBody)
			}, func(c *gin.Context, requestBody *models.PostCategoryBody) (*models.Category, error) {
				return services.CreateCategory(requestBody)
			})
		}))
		routes.PUT("/:id", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Put(ctx, func(c *gin.Context, requestBody *models.PutCategoryBody) error {
				return validators.ValidatePutCategoryBody(requestBody)
			}, func(c *gin.Context, ID uint, requestBody *models.PutCategoryBody) (*models.Category, error) {
				return services.UpdateCategory(ID, requestBody)
			},
			)
		}))
	}
}
