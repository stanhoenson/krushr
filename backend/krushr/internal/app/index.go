package app

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	env.InitializeEnvironment()

	r := gin.Default()
	validators.InitializeValidators()
	r.Use(middleware.Authorization())
	handlers.InitializeHandlers(r)
	database.InitializeDatabase(env.DatabaseName)

	r.Run(env.Address)
}
