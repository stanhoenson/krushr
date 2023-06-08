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
		routes.GET("", GetAllDefault[models.Detail])
		routes.DELETE("/:id", wrappers.RoleWrapper([]string{constants.AdminRoleName}, DeleteByIDDefault[models.Detail]))
	}
}
