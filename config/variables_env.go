package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	ServerPort string
	Secret     string
)

func InitEnvVariables(file string) {
	if file == "" {
		file = ".env"
	}
	err := godotenv.Load(file)
	if err != nil {
		log.Fatal(err)
	}
	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort = os.Getenv("DB_PORT")
	ServerPort = os.Getenv("SERVER_PORT")
	Secret = os.Getenv("SECRET")
}
