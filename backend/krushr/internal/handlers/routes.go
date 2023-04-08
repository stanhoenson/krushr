package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func putRoute(c *gin.Context) {
	var updatedRoute models.Route

	if err := c.BindJSON(&updatedRoute); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePutRoute(&updatedRoute); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.UpdateEntity(&updatedRoute)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating route"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedRoute)
}

func deleteRouteByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedRoute, err := services.DeleteEntityByID[models.Route](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting route"})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedRoute)
}

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

	c.IndentedJSON(http.StatusCreated, newRoute)
}

func RoutesRoutes(r *gin.Engine) {
	routes := r.Group("/routes")
	{
		routes.GET("", getRoutes)
		routes.POST("", postRoute)
		routes.PUT("", putRoute)
		routes.GET("/:id", getRouteByID)
		routes.DELETE("/:id", deleteRouteByID)
	}
}
