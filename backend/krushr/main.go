package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/routes", GetRoutes)
	r.GET("/routes/:id", GetRouteByID)
	r.POST("/routes", PostRoutes)

	r.Run("localhost:8080")
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
