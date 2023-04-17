package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	JWTSecret       string
	DatabaseName    string
	Address         string
	DefaultRoleID   uint
	FileStoragePath string
)

func InitializeEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		panicMessage(JWTSecret)
	}

	DatabaseName = os.Getenv("DATABASE_NAME")
	if DatabaseName == "" {
		panicMessage(DatabaseName)
	}

	Address = os.Getenv("ADDRESS")
	if Address == "" {
		panicMessage(Address)
	}
	FileStoragePath = os.Getenv("FILE_STORAGE_PATH")
	if FileStoragePath == "" {
		panicMessage(FileStoragePath)
	}

	roleIDString := os.Getenv("DEFAULT_ROLE_ID")
	u64, err := strconv.ParseUint(roleIDString, 10, 64)
	if err != nil {
		panicMessage(roleIDString)
	}
	DefaultRoleID = uint(u64)
}

func panicMessage(variable string) {
	message := "failed to get " + variable + " environment variable"
	panic(message)
}
