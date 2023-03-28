package app

import (
	"github.com/stanhoenson/krushr/internal/config"
	"github.com/stanhoenson/krushr/internal/database"
	"github.com/stanhoenson/krushr/internal/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	config *config.Config
}

func CreateApp() (*App, error) {
	newConfig, err := config.NewConfig()
	if err != nil {
		panic("couldn't create config")
	}
	app := App{
		config: newConfig,
	}

	return &app, nil
}

func Initialize(app *App) {
	r := gin.Default()
	handlers.InitializeHandlers(r)
	database.InitializeDatabase(&app.config.Database)

	r.Run()
}
