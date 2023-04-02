package handlers

import (
	"net/http"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func GetEntries(c *gin.Context) {

	entries, err := services.GetEntites[models.Entry]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving entries"})
		return
	}
	c.IndentedJSON(http.StatusOK, entries)
}
func EntriesRoutes(r *gin.Engine) {
	routes := r.Group("/entries")
	{
		routes.GET("", GetEntries)
	}
}
