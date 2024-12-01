package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser       string
	PostgresPassword   string
	PostgresDB         string
	PostgresHost       string
	PostgresPort       string
	JWTSecret          string
	JwtExpirationHours string
	GscBucketName      string
	GscFolderPeople    string
	GscFolderMask      string
	GscFolderClothes   string
	GscKeyFile         string
	HostServer         string
	PortServer         string
	ModelGenAPI        string
	GscStoreImage      string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment variables")
	}

	AppConfig = Config{
		PostgresUser:       getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword:   getEnv("POSTGRES_PASSWORD", "password"),
		PostgresDB:         getEnv("POSTGRES_DB", "postgres_db"),
		PostgresHost:       getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:       getEnv("POSTGRES_PORT", "5432"),
		JWTSecret:          getEnv("JWT_SECRET", "default_jwt_secret"),
		JwtExpirationHours: getEnv("JWT_EXPIRATION_HOURS", "72"),
		GscBucketName:      getEnv("GSC_BUCKET_NAME", ""),
		GscFolderPeople:    getEnv("GSC_FOLDER_PEOPLE", ""),
		GscFolderMask:      getEnv("GSC_FOLDER_POSH", ""),
		GscFolderClothes:   getEnv("GSC_FOLDER_CLOTHES", ""),
		GscKeyFile:         getEnv("GSC_KEY_FILE", ""),
		HostServer:         getEnv("HOST_SERVER", "localhost"),
		PortServer:         getEnv("PORT_SERVER", "8080"),
		ModelGenAPI:        getEnv("MODEL_GEN_API", ""),
		GscStoreImage:      getEnv("GSC_STORE_IMAGE", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
