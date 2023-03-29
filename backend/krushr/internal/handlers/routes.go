package handlers

import (
	"fmt"
	"net/http"

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
	routes := services.GetRoutes()
	fmt.Println(routes)
	c.IndentedJSON(http.StatusOK, routes)
}

func getRouteByID(c *gin.Context) {
	// id := c.Param("id")

	// for _, r := range routes {
	// 	if r.ID == id {
	// 		c.IndentedJSON(http.StatusOK, r)
	// 		return
	// 	}
	// }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "route not found"})
}

func postRoutes(c *gin.Context) {
	var newRoute models.Route

	if err := c.BindJSON(&newRoute); err != nil {
		return
	}

	// routes = append(routes, newRoute)
	c.IndentedJSON(http.StatusCreated, newRoute)
}

func Routes(r *gin.Engine) {
	routes := r.Group("/routes")
	{
		routes.GET("", getRoutes)
		routes.POST("", postRoutes)
		routes.GET("/:id", getRouteByID)
	}
}
