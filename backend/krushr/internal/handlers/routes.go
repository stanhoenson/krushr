package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/stanhoenson/krushr/internal/wrappers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PutRoute(c *gin.Context) {
	var putRouteBody models.PutRouteBody

	if err := c.BindJSON(&putRouteBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePutRoute(&putRouteBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user in context"})
		return
	}

	updatedRoute, err := services.UpdateRoute(&putRouteBody, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating route"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedRoute)
}

func DeleteRouteByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	user, err := utils.GetUserFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No user in context"})
		return
	}

	deletedRoute, err := services.DeleteRouteByIDAndAuthenticatedUser(ID, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting route"})
		return
	}

	c.JSON(http.StatusOK, deletedRoute)
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

	createdRoute, err := wrappers.WithTransaction(database.Db, func(tx *gorm.DB) (*models.Route, error) {
		return services.CreateRoute(&postRouteBody, user, tx)

	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating route" + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, createdRoute)
}

func RegisterRouteRoutes(r *gin.Engine) {
	routes := r.Group("/routes")
	{
		routes.GET("", getRoutes)
		routes.POST("", wrappers.RoleWrapper(constants.Roles, postRoute))
		routes.PUT("", wrappers.RoleWrapper(constants.Roles, PutRoute))
		routes.GET("/:id", GetRouteByID)
		routes.DELETE("/:id", wrappers.RoleWrapper(constants.Roles, DeleteRouteByID))
	}
}
