package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	JWTSecret         string
	DatabaseName      string
	Address           string
	DefaultRoleID     uint
	FileStorageFolder string
	DataFolder        string
	ApiUrl            string
	Domain            string
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
	FileStorageFolder = os.Getenv("FILE_STORAGE_FOLDER")
	if FileStorageFolder == "" {
		panicMessage(FileStorageFolder)
	}

	DataFolder = os.Getenv("DATA_FOLDER")
	if DataFolder == "" {
		panicMessage(DataFolder)
	}
	ApiUrl = os.Getenv("API_URL")
	if ApiUrl == "" {
		panicMessage(ApiUrl)
	}
	Domain = os.Getenv("DOMAIN")
	if Domain == "" {
		panicMessage(Domain)
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
