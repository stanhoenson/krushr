package app

import (
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	env.Initialize()

	r := gin.Default()
	validators.InitializeValidators()
	handlers.InitializeHandlers(r)
	database.InitializeDatabase()

	r.Run(env.Address)
}
