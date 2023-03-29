package database

import (
	"os"

	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase() {
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		panic("failed to get environment variable")
	}

	db, err := gorm.Open(sqlite.Open(databaseName))

	// TODO look at ways to do this nicely
	db.AutoMigrate(&models.Category{}, &models.Entry{}, &models.PointOfInterest{}, &models.Role{}, &models.Route{}, &models.Status{}, &models.Type{}, &models.User{})

	if err != nil {
		panic("failed to initialize database")
	}

	Db = db
}
