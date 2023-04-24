package handlers

import "github.com/gin-gonic/gin"

func RegisterLegacyRoutes(r *gin.Engine) {
	routes := r.Group("/legacy")
	{
		routes.GET("/route")
		routes.GET("/route/:id")
		routes.GET("/routeimage/:id")
		routes.GET("/menu")
	}
}
