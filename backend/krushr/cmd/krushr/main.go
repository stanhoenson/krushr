package main

import (
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	handlers.InitHandlers(r)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
