package handlers

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

var userPostOptions = PostOptions[models.User, models.PostUserBody]{
	ValidationFunction: func(requestBody *models.PostUserBody) error {
		return nil
	},
	CreateFunction: func(c *gin.Context, requestBody *models.PostUserBody) (*models.User, error) {
		return services.CreateUser(requestBody, database.Db)
	},
}

var userPutOptions = PutOptions[models.User, models.PutUserBody]{
	ValidationFunction: func(requestBody *models.PutUserBody) error {
		return nil
	},
	UpdateFunction: func(c *gin.Context, ID uint, requestBody *models.PutUserBody) (*models.User, error) {
		return services.UpdateUser(ID, requestBody, database.Db)
	},
}

func RegisterUserRoutes(r *gin.Engine) {
	routes := r.Group("/users")
	{
		routes.GET("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			GetAll[models.User](ctx)
		}))
		routes.POST("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			Post(ctx, func(po *PostOptions[models.User, models.PostUserBody]) {
				po = &userPostOptions
			})
		}))
		routes.PUT("", wrappers.RoleWrapper([]string{constants.AdminRoleName}, func(ctx *gin.Context) {
			Put(ctx, func(po *PutOptions[models.User, models.PutUserBody]) {
				po = &userPutOptions
			})
		}))
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByID[models.User]))
	}
}
