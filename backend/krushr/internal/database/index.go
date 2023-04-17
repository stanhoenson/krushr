package database

import (
	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase(database string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(database))

	// TODO look at ways to do this nicely, only in development
	db.AutoMigrate(&models.Route{}, &models.Image{}, &models.Detail{}, &models.Link{}, &models.Category{}, &models.Status{}, &models.PointOfInterest{}, &models.User{}, &models.Role{})

	if err != nil {
		panic("failed to initialize database")
	}

	Db = db
	populateDatabase()
	return Db
}

func populateDatabase() {
	Db.Create(&models.Role{Name: constants.AdminRoleName})
	Db.Create(&models.Role{Name: constants.CreatorRoleName})
}
