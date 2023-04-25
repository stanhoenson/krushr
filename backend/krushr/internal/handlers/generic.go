package handlers

import (
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/gin-gonic/gin"
)

func DeleteByIDDefault[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	DeleteByID(c, func(c *gin.Context, ID uint) (*EntityType, error) {
		return services.DeleteEntityByID[EntityType](ID)
	})
}
func DeleteByID[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context, deleteFunction func(c *gin.Context, ID uint) (*EntityType, error)) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedEntity, err := deleteFunction(c, ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting " + utils.GetTypeString(deletedEntity)})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedEntity)
}

func GetAllDefault[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	GetAll(c, func(c *gin.Context) (*[]EntityType, error) {
		return services.GetEntities[EntityType]()
	})
}
func GetAll[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context, getFunction func(c *gin.Context) (*[]EntityType, error)) {
	entities, err := getFunction(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving " + utils.GetTypeString(entities)})
		return
	}

	c.IndentedJSON(http.StatusOK, entities)
}
func GetByIDDefault[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	GetByID(c, func(c *gin.Context, ID uint) (*EntityType, error) {
		return services.GetEntityByID[EntityType](ID)
	})
}

func GetByID[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context, getByIDFunction func(c *gin.Context, ID uint) (*EntityType, error)) {

	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	entity, err := getByIDFunction(c, ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving " + utils.GetTypeString(entity)})
		return
	}

	c.IndentedJSON(http.StatusOK, entity)
}

func Post[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest, RequestBodyType any](c *gin.Context, validationFunction func(c *gin.Context, requestBody *RequestBodyType) error, createFunction func(c *gin.Context, requestBody *RequestBodyType) (*EntityType, error)) {

	var requestBody RequestBodyType

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validationFunction(c, &requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEntity, err := createFunction(c, &requestBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating " + utils.GetTypeString(createdEntity) + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, createdEntity)
}

func Put[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest, RequestBodyType any](c *gin.Context, validationFunction func(c *gin.Context, requestBody *RequestBodyType) error, updateFunction func(c *gin.Context, ID uint, requestBody *RequestBodyType) (*EntityType, error)) {

	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	var requestBody RequestBodyType

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validationFunction(c, &requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEntity, err := updateFunction(c, ID, &requestBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating " + utils.GetTypeString(updatedEntity)})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedEntity)
}
