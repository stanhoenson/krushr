package app

import (
	"flag"

	"github.com/stanhoenson/krushr/internal/cors"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/middleware"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	env.InitializeEnvironment("./.env")

	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	cors.InitializeCors(r)
	validators.InitializeValidators()
	r.Use(middleware.Authorization())
	handlers.RegisterHandlers(r)
	database.InitializeDatabase(env.DatabaseName, env.DataFolder)

	r.Run(env.Address)
}
