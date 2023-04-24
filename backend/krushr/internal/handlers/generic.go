package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/services"
	"github.com/stanhoenson/krushr/internal/utils"
	"github.com/gin-gonic/gin"
)

type DeleteByIDOptions[T any] struct {
	DeleteFunction func(c *gin.Context, ID uint) (*T, error)
}

type DeleteByIDOption[T any] func(*DeleteByIDOptions[T])

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting " + utils.GetTypeString(deletedEntity)})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedEntity)
}

func DeleteByIDCool[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context, deleteFunction func(c *gin.Context, ID uint) (*T, error)) {
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

func DeleteByIDCooler[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context, setters ...DeleteByIDOption[T]) {
	deleteByIDOptions := &DeleteByIDOptions[T]{
		DeleteFunction: func(c *gin.Context, ID uint) (*T, error) {
			return services.DeleteEntityByID[T](ID)
		},
	}

	for _, setter := range setters {
		setter(deleteByIDOptions)
	}

	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	deletedEntity, err := deleteByIDOptions.DeleteFunction(c, ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error deleting " + utils.GetTypeString(deletedEntity)})
		return
	}

	c.IndentedJSON(http.StatusOK, deletedEntity)
}

type GetAllOptions[T any] struct {
	GetFunction func(c *gin.Context) (*[]T, error)
}

type GetAllOption[T any] func(*GetAllOptions[T])

func GetAll[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context, setters ...GetAllOption[T]) {
	getAllOptions := &GetAllOptions[T]{
		GetFunction: func(c *gin.Context) (*[]T, error) {
			return services.GetEntities[T]()
		},
	}

	for _, setter := range setters {
		setter(getAllOptions)
	}

	entities, err := getAllOptions.GetFunction(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving " + utils.GetTypeString(entities)})
		return
	}

	c.IndentedJSON(http.StatusOK, entities)
}

type GetByIDOptions[T any] struct {
	GetFunction func(c *gin.Context, ID uint) (*T, error)
}

type GetByIDOption[T any] func(*GetByIDOptions[T])

func GetByID[T models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest](c *gin.Context, setters ...GetByIDOption[T]) {
	getByIDOptions := &GetByIDOptions[T]{
		GetFunction: func(c *gin.Context, ID uint) (*T, error) {
			return services.GetEntityByID[T](ID)
		},
	}

	for _, setter := range setters {
		setter(getByIDOptions)
	}

	id := c.Param("id")

	u64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	ID := uint(u64)

	entity, err := getByIDOptions.GetFunction(c, ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving " + utils.GetTypeString(entity)})
		return
	}

	c.IndentedJSON(http.StatusOK, entity)
}

type PostOptions[EntityType any, RequestBodyType any] struct {
	ValidationFunction func(requestBody *RequestBodyType) error
	CreateFunction     func(c *gin.Context, requestBody *RequestBodyType) (*EntityType, error)
}

type PostOption[EntityType any, RequestBodyType any] func(*PostOptions[EntityType, RequestBodyType])

func Post[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest, RequestBodyType any](c *gin.Context, setters ...PostOption[EntityType, RequestBodyType]) {
	postOptions := &PostOptions[EntityType, RequestBodyType]{
		ValidationFunction: func(requestBody *RequestBodyType) error { return nil },
		CreateFunction: func(c *gin.Context, requestBody *RequestBodyType) (*EntityType, error) {
			return nil, fmt.Errorf("Not implemented")
		},
	}
	for _, setter := range setters {
		setter(postOptions)
	}

	var requestBody RequestBodyType

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := postOptions.ValidationFunction(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEntity, err := postOptions.CreateFunction(c, &requestBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error creating " + utils.GetTypeString(createdEntity) + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, createdEntity)
}

type PutOptions[EntityType any, RequestBodyType any] struct {
	ValidationFunction func(requestBody *RequestBodyType) error
	UpdateFunction     func(c *gin.Context, ID uint, requestBody *RequestBodyType) (*EntityType, error)
}

type PutOption[EntityType any, RequestBodyType any] func(*PutOptions[EntityType, RequestBodyType])

func Put[EntityType models.Route | models.Image | models.Detail | models.Link | models.Category | models.Status | models.PointOfInterest | models.User | models.Role | models.RoutesPointsOfInterest, RequestBodyType any](c *gin.Context, setters ...PutOption[EntityType, RequestBodyType]) {
	putOptions := &PutOptions[EntityType, RequestBodyType]{
		ValidationFunction: func(requestBody *RequestBodyType) error {
			return nil
		},
		UpdateFunction: func(c *gin.Context, ID uint, requestBody *RequestBodyType) (*EntityType, error) {
			return nil, fmt.Errorf("Not implemented")
		},
	}
	for _, setter := range setters {
		setter(putOptions)
	}

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

	if err := putOptions.ValidationFunction(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEntity, err := putOptions.UpdateFunction(c, ID, &requestBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error updating " + utils.GetTypeString(updatedEntity)})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedEntity)
}
