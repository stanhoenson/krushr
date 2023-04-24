package handlers

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func RegisterDetailRoutes(r *gin.Engine) {
	routes := r.Group("/details")
	{
		routes.GET("", func(ctx *gin.Context) {
			GetAll[models.Detail](ctx)
		})
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByID[models.Detail]))
	}
}
