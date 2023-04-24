package handlers

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func RegisterLinkRoutes(r *gin.Engine) {
	routes := r.Group("/links")
	{
		routes.GET("", func(ctx *gin.Context) {
			GetAll[models.Link](ctx)
		})
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByID[models.Link]))
	}
}
