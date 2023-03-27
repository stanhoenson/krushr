package database

import (
	"github.com/stanhoenson/krushr/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeDatabase(databaseConfig *config.DatabaseConfig) (*gorm.DB, error) {

	var db, err = gorm.Open(sqlite.Open(databaseConfig.Name))

	return db, err

}
