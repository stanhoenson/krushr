package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DatabaseName string
var JwtSecret string
var Address string
var DefaultRoleID uint

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JwtSecret = os.Getenv("JWT_SECRET")
	if JwtSecret == "" {
		panic("failed to get JWT_SECRET environment variable")
	}

	DatabaseName = os.Getenv("DATABASE_NAME")
	if DatabaseName == "" {
		panic("failed to get DATABASE_NAME environment variable")
	}

	Address = os.Getenv("ADDRESS")
	if Address == "" {
		panic("failed to get ADDRESS environment variable")
	}
	roleIdString := os.Getenv("DEFAULT_ROLE_ID")

	u64, err := strconv.ParseUint(roleIdString, 10, 64)
	if err != nil {
		panic("failed to get DEFAULT_ROLE_ID environment variable")
	}
	DefaultRoleID = uint(u64)

}
