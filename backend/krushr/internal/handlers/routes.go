package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

// var routes = []models.Route{
// 	{ID: "0", Title: "Bernard's Route", StatusID: "0", UserID: "2"},
// 	{ID: "1", Title: "Het Pad der 7 Zonden", StatusID: "0", UserID: "1"},
// 	{ID: "2", Title: "Reflecties en Introspecties", StatusID: "1", UserID: "2"},
// }

func getRoutes(c *gin.Context) {
	routes, err := services.GetRoutes()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving routes"})
		return
	}
	c.IndentedJSON(http.StatusOK, routes)
}

func getRouteByID(c *gin.Context) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	route, err := services.GetRouteByID(ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving routes"})
		return
	}

	c.IndentedJSON(http.StatusOK, route)
}

func postRoutes(c *gin.Context) {
	var newRoute models.Route

	if err := c.BindJSON(&newRoute); err != nil {
		return
	}

	// routes = append(routes, newRoute)
	c.IndentedJSON(http.StatusCreated, newRoute)
}

func RoutesRoutes(r *gin.Engine) {
	routes := r.Group("/routes")
	{
		routes.GET("", getRoutes)
		routes.POST("", postRoutes)
		routes.GET("/:id", getRouteByID)
	}
}
