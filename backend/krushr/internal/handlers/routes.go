package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
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
	id := c.Param("id")

	// Convert string to uint
	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	route, err := services.GetEntity[models.Route](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving routes"})
		return
	}

	c.IndentedJSON(http.StatusOK, route)
}
func postRoute(c *gin.Context) {
	var newRoute models.Route

	if err := c.BindJSON(&newRoute); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostRoute(&newRoute); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRoute, err := services.CreateEntity(&newRoute)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating route"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdRoute)
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
		routes.POST("", postRoute)
		routes.GET("/:id", getRouteByID)
	}
}
