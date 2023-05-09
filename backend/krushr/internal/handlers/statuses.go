package handlers

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
)

func RegisterStatusRoutes(r *gin.Engine) {
	routes := r.Group("/statuses")
	{
		routes.GET("", GetAllDefault[models.Status])
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByIDDefault[models.Status]))
	}
}
