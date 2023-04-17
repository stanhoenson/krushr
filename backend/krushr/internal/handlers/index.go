package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitializeHandlers(r *gin.Engine) {
	RegisterRouteRoutes(r)
	RegisterPointOfInterestRoutes(r)
	RegisterUserRoutes(r)
	RegisterCategoryRoutes(r)
	RegisterAuthenticationRoutes(r)
	RegisterImageRoutes(r)
}
