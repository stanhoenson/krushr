package handlers

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

var categoryPostOptions = PostOptions[models.Category, models.PostCategoryBody]{
	ValidationFunction: func(requestBody *models.PostCategoryBody) error {
		return validators.ValidatePostCategoryBody(requestBody)
	},
	CreateFunction: func(c *gin.Context, requestBody *models.PostCategoryBody) (*models.Category, error) {
		return services.CreateCategory(requestBody)
	},
}
var categoryPutOptions = PutOptions[models.Category, models.PutCategoryBody]{

	ValidationFunction: func(requestBody *models.PutCategoryBody) error {
		return validators.ValidatePutCategoryBody(requestBody)
	},
	UpdateFunction: func(c *gin.Context, ID uint, requestBody *models.PutCategoryBody) (*models.Category, error) {
		return services.UpdateCategory(ID, requestBody)
	},
}

func RegisterCategoryRoutes(r *gin.Engine) {
	routes := r.Group("/categories")
	{
		routes.GET("", func(ctx *gin.Context) {
			GetAll[models.Category](ctx)
		})
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByID[models.Category]))
		routes.POST("", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Post(ctx, func(po *PostOptions[models.Category, models.PostCategoryBody]) {
				po = &categoryPostOptions
			})

		}))
		routes.PUT("", wrappers.RoleWrapper(constants.Roles, func(ctx *gin.Context) {
			Put(ctx, func(po *PutOptions[models.Category, models.PutCategoryBody]) {
				po = &categoryPutOptions
			})
		}))
	}
}
