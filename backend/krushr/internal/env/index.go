package env

import (
	"log"
	"os"
	"strconv"
	"strings"

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
	AdminPassword     string
	AllowedOrigins    []string
)

func InitializeEnvironment() {
	loadEnv()

	JWTSecret = getEnvVariable("JWT_SECRET")
	DatabaseName = getEnvVariable("DATABASE_NAME")
	Address = getEnvVariable("ADDRESS")
	FileStorageFolder = getEnvVariable("FILE_STORAGE_FOLDER")
	DataFolder = getEnvVariable("DATA_FOLDER")
	ApiUrl = getEnvVariable("API_URL")
	Domain = getEnvVariable("DOMAIN")
	AdminPassword = getEnvVariable("ADMIN_PASSWORD")
	AllowedOrigins = strings.Split(getEnvVariable("ALLOWED_ORIGINS"),",")

	DefaultRoleID = getUintEnvVariable("DEFAULT_ROLE_ID")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnvVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panicMessage(key)
	}
	return value
}

func getUintEnvVariable(key string) uint {
	valueString := getEnvVariable(key)
	value, err := strconv.ParseUint(valueString, 10, 64)
	if err != nil {
		panicMessage(key)
	}
	return uint(value)
}

func panicMessage(variable string) {
	message := "failed to get " + variable + " environment variable"
	panic(message)
}
