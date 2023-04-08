package app

import (
	"log"
	"os"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/stanhoenson/krushr/internal/validators"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	validators.InitializeValidators()
	handlers.InitializeHandlers(r)
	database.InitializeDatabase()

	address := os.Getenv("ADDRESS")
	if address == "" {
		panic("failed to load environment variable")
	}

	r.Run(address)
}
