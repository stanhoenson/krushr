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
	DeleteByID(c, func(c *gin.Context, ID uint) (uint, error) {
		return services.DeleteEntityByID[EntityType](ID)
	})
}

func DeleteByID(c *gin.Context, deleteFunction func(c *gin.Context, ID uint) (uint, error)) {
	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedEntityID, err := deleteFunction(c, ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting entity"})
		return
	}

	c.JSON(http.StatusOK, deletedEntityID)
}

func GetAllDefault[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	GetAll(c, func(c *gin.Context) (*[]EntityType, error) {
		return services.GetEntities[EntityType]()
	})
}

func GetAll[EntityType any](c *gin.Context, getFunction func(c *gin.Context) (*[]EntityType, error)) {
	entities, err := getFunction(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving " + utils.GetTypeString(entities) + err.Error()})
		return
	}

	c.JSON(http.StatusOK, entities)
}

func GetByIDDefault[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context) {
	GetByID(c, func(c *gin.Context, ID uint) (*EntityType, error) {
		return services.GetEntityByID[EntityType](ID)
	})
}

func GetByID[EntityType any](c *gin.Context, getByIDFunction func(c *gin.Context, ID uint) (*EntityType, error)) {
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

	c.JSON(http.StatusOK, entity)
}

func Post[EntityType any, RequestBodyType any](c *gin.Context, validationFunction func(c *gin.Context, requestBody *RequestBodyType) error, createFunction func(c *gin.Context, requestBody *RequestBodyType) (*EntityType, error)) {
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

	c.JSON(http.StatusOK, createdEntity)
}

func Put[EntityType any, RequestBodyType any](c *gin.Context, validationFunction func(c *gin.Context, requestBody *RequestBodyType) error, updateFunction func(c *gin.Context, ID uint, requestBody *RequestBodyType) (*EntityType, error)) {
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating " + utils.GetTypeString(updatedEntity) + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEntity)
}
