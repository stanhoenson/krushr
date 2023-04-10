package database

import (
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase(database string) {
	db, err := gorm.Open(sqlite.Open(database))

	// TODO look at ways to do this nicely
	db.AutoMigrate(&models.Category{}, &models.Entry{}, &models.PointOfInterest{}, &models.Role{}, &models.Route{}, &models.Status{}, &models.Type{}, &models.User{})

	if err != nil {
		panic("failed to initialize database")
	}

	Db = db
}
