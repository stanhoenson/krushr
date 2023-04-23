package handlers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"

	"github.com/stanhoenson/krushr/internal/services"
	"github.com/gin-gonic/gin"
)

func DeleteByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedEntity, err := services.DeleteEntityByID[T](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting " + reflect.TypeOf(deletedEntity).String()})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedEntity)
}

func GetAll[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	entities, err := services.GetEntities[T]()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving" + reflect.TypeOf(entities).String()})
		return
	}

	c.IndentedJSON(http.StatusOK, entities)
}

func GetByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	entity, err := services.GetEntity[T](ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving" + reflect.TypeOf(entity).String()})
		return
	}

	c.IndentedJSON(http.StatusOK, entity)
}

func Post[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest, RequestBodyType any](c *gin.Context, validationFunction func(requestBody *RequestBodyType) error, createFunction func(c *gin.Context, requestBody *RequestBodyType) (*EntityType, error)) {
	var requestBody RequestBodyType

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validationFunction(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEntity, err := createFunction(c, &requestBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating " + reflect.TypeOf(createdEntity).String()})
		return
	}

	c.IndentedJSON(http.StatusOK, createdEntity)
}

func Put[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest, RequestBodyType any](c *gin.Context, validationFunction func(requestBody *RequestBodyType) error, updateFunction func(c *gin.Context, ID uint, requestBody *RequestBodyType) (*EntityType, error)) {
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

	if err := validationFunction(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEntity, err := updateFunction(c, ID, &requestBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating " + reflect.TypeOf(updatedEntity).String()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedEntity)
}
