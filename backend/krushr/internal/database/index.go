package database

import (
	"github.com/stanhoenson/krushr/internal/config"
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase(databaseConfig *config.DatabaseConfig) {
	db, err := gorm.Open(sqlite.Open(databaseConfig.Name))

	// TODO look at ways to do this nicely
	db.AutoMigrate(&models.Route{})

	if err != nil {
		panic("failed to initialize database")
	}

	Db = db
}
