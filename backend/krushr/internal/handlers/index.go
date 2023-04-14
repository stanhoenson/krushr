package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitializeHandlers(r *gin.Engine) {
	RoutesRoutes(r)
	PointsOfInterestRoutes(r)
	UsersRoutes(r)
	CategoriesRoutes(r)
	AuthenticationRoutes(r)
}
