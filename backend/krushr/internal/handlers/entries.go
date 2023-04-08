package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func putEntry(c *gin.Context) {
	var updatedEntry models.Entry

	if err := c.BindJSON(&updatedEntry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePutEntry(&updatedEntry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.UpdateEntity(&updatedEntry)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating entry"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedEntry)
}

func deleteEntryByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedEntry, err := services.DeleteEntityByID[models.Entry](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting entry"})
		return
	}
	c.IndentedJSON(http.StatusOK, deletedEntry)
}

func getEntries(c *gin.Context) {
	entries, err := services.GetEntites[models.Entry]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving entries"})
		return
	}

	c.IndentedJSON(http.StatusOK, entries)
}

func getEntryByID(c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	entry, err := services.GetEntity[models.Entry](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving entry"})
		return
	}

	c.IndentedJSON(http.StatusOK, entry)
}

func postEntry(c *gin.Context) {
	var newEntry models.Entry

	if err := c.BindJSON(&newEntry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validators.ValidatePostEntry(&newEntry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEntry, err := services.CreateEntity(&newEntry)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating entry"})
		return
	}

	c.IndentedJSON(http.StatusOK, createdEntry)
}

func EntriesRoutes(r *gin.Engine) {
	routes := r.Group("/entries")
	{
		routes.GET("", getEntries)
		routes.GET("/:id", getEntryByID)
		routes.POST("", postEntry)
		routes.PUT("", putEntry)
		routes.DELETE("/:id", deleteEntryByID)
	}
}
