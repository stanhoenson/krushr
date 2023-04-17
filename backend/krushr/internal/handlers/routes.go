package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func PutRoute(c *gin.Context) {
	var updatedRoute models.PutRouteBody

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

func DeleteRouteByID(c *gin.Context) {
	hasRoles := utils.HasRole(c, constants.Roles)
	if !hasRoles {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No role"})
		return
	}
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	user, err := utils.GetUserFromContext(c)
	fmt.Println(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user in context"})
		return
	}

	deletedRoute, err := services.DeleteRouteByIDAndAuthenticatedUser(ID, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting route"})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedRoute)
}

func getRoutes(c *gin.Context) {
	routes, err := services.GetEntities[models.Route]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving routes"})
		return
	}

	c.IndentedJSON(http.StatusOK, routes)
}

func GetRouteByID(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving route"})
		return
	}

	c.IndentedJSON(http.StatusOK, route)
}

func postRoute(c *gin.Context) {
	var postRouteBody models.PostRouteBody

	if err := c.BindJSON(&postRouteBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostRouteBody(&postRouteBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user in context"})
		return
	}

	createdRoute, err := services.CreateRoute(&postRouteBody, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating route"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdRoute)
}

func RegisterRouteRoutes(r *gin.Engine) {
	routes := r.Group("/routes")
	{
		routes.GET("", getRoutes)
		routes.POST("", postRoute)
		routes.PUT("", PutRoute)
		routes.GET("/:id", GetRouteByID)
		routes.DELETE("/:id", DeleteRouteByID)
	}
}
