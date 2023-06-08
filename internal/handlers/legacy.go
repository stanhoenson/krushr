package handlers

import (
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterLegacyRoutes(r *gin.Engine) {
	routes := r.Group("/legacy")
	{
		routes.GET("/route", func(ctx *gin.Context) {
			GetAll(ctx, func(c *gin.Context) (*[]models.LegacyRoute, error) {
				return services.GetLegacyRoutes()
			})
		})
		routes.GET("/route/:id", func(ctx *gin.Context) {
			GetByID(ctx, func(c *gin.Context, ID uint) (*models.LegacyRoute, error) {
				return services.GetLegacyRouteByID(ID)
			})
		})
		routes.GET("/routeimage/:id", getImageDataByID)
		routes.GET("/menu", func(ctx *gin.Context) {
			GetAll(ctx, func(c *gin.Context) (*[]models.LegacyMenu, error) {
				return services.GetLegacyMenus()
			},
			)
		})
	}
}
