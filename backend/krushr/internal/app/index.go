package app

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	env.InitializeEnvironment()

	// Returns an engine with a Logger and Recovery middleware already attached
	r := gin.Default()
	r.Use(cors.Default())
	validators.InitializeValidators()
	r.Use(middleware.Authorization())
	handlers.RegisterHandlers(r)
	database.InitializeDatabase(env.DatabaseName, env.DataFolder)

	r.Run(env.Address)
}
