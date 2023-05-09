package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	RegisterLinkRoutes(r)
	RegisterDetailRoutes(r)
	RegisterRouteRoutes(r)
	RegisterPointOfInterestRoutes(r)
	RegisterUserRoutes(r)
	RegisterCategoryRoutes(r)
	RegisterAuthenticationRoutes(r)
	RegisterImageRoutes(r)
	RegisterLegacyRoutes(r)
	RegisterStatusRoutes(r)
}
