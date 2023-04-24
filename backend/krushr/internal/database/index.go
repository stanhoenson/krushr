package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase(databaseName, folderName string) *gorm.DB {

	// Create the directory if it doesn't exist
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal("failed to create directory")
	}

	db, err := gorm.Open(sqlite.Open(filepath.Join(folderName, databaseName)))
	if err != nil {
		log.Fatal("failed to initialize database")
	}
	// TODO look at ways to do this nicely, only in development
	err = db.AutoMigrate(&models.Route{}, &models.Image{}, &models.Detail{}, &models.Link{}, &models.Category{}, &models.Status{}, &models.PointOfInterest{}, &models.User{}, &models.Role{}, &models.RoutesPointsOfInterest{})

	if err != nil {
		log.Fatal("failed to auto migrate models")
	}
	// declare some custom things
	err = db.SetupJoinTable(&models.Route{}, "PointsOfInterest", &models.RoutesPointsOfInterest{})
	if err != nil {
		log.Fatal("failed to setup join table")
	}
	err = db.SetupJoinTable(&models.PointOfInterest{}, "Routes", &models.RoutesPointsOfInterest{})
	if err != nil {
		log.Fatal("failed to setup join table")
	}

	Db = db
	populateDatabase()
	return Db
}

func populateDatabase() {
	Db.Create(&models.Role{Name: constants.AdminRoleName})
	Db.Create(&models.Role{Name: constants.CreatorRoleName})

	//statuses
	for _, statusName := range constants.Statuses {
		Db.Create(&models.Status{Name: statusName})
	}
}
