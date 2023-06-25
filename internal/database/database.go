package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/stanhoenson/krushr/internal/constants"
	"github.com/stanhoenson/krushr/internal/env"
	"github.com/stanhoenson/krushr/internal/models"
	"github.com/stanhoenson/krushr/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase(databaseName, folderName string, automigrate bool) *gorm.DB {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal("failed to create directory")
	}

	db, err := gorm.Open(sqlite.Open(filepath.Join(folderName, databaseName)))
	if err != nil {
		log.Fatal("failed to initialize database")
	}

	if automigrate {
		err = db.AutoMigrate(&models.Route{}, &models.Image{}, &models.Detail{}, &models.Link{}, &models.Category{}, &models.Status{}, &models.PointOfInterest{}, &models.User{}, &models.Role{}, &models.RoutesPointsOfInterest{})
		if err != nil {
			log.Fatal("failed to auto migrate models")
		}
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
	result := Db.Save(&models.Role{ID: 1, Name: constants.AdminRoleName})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	result = Db.Save(&models.Role{ID: 2, Name: constants.CreatorRoleName})
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// We have to hash here because we hash on the client as well
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(utils.Sha256(env.AdminPassword+"b6570a1b5af373d269be27b8c1244650")), bcrypt.DefaultCost)
	Db.Save(&models.User{ID: 1, Email: "admin@admin.com", Password: string(passwordBytes), RoleID: 1})
	if err != nil {
		log.Fatal(err)
	}

	// TODO hmmmmm, should be better
	categories := []string{"Default"}
	for index, category := range categories {
		result = Db.Save(&models.Category{ID: uint(index + 1), Name: category})
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}

	// statuses
	for index, statusName := range constants.Statuses {
		result = Db.Save(&models.Status{ID: uint(index + 1), Name: statusName})
		if result.Error != nil {
			log.Fatal(result.Error)
		}
	}
}
