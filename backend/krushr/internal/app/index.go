package app

import (
	"log"
	"os"

	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func CreateApp() (*App, error) {
// 	newConfig, err := config.NewConfig()
// 	if err != nil {
// 		panic("couldn't create config")
// 	}
// 	app := App{
// 		config: newConfig,
// 	}

// 	return &app, nil
// }

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	handlers.InitializeHandlers(r)
	database.InitializeDatabase()

	address := os.Getenv("ADDRESS")
	if address == "" {
		panic("failed to load environment variable")
	}
	r.Run(address)
}
