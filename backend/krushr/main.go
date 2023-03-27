package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/routes", GetRoutes)
	router.GET("/routes/:id", GetRouteByID)
	router.POST("/routes", PostRoutes)

	router.Run("localhost:8080")
}
