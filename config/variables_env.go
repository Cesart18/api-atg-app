package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUrl      string
	ServerPort string
	Secret     string
	AdminKey   string
)

func InitEnvVariables(file string) {
	if file == "" {
		file = ".env"
	}
	err := godotenv.Load(file)
	if err != nil {
		log.Fatal(err)
	}
	DBUrl = os.Getenv("DATABASE_URL")
	if DBUrl == "" {
		log.Fatal("DATABASE_URL not set")
	}

	ServerPort = os.Getenv("SERVER_PORT")
	if ServerPort == "" {
		ServerPort = "8080"
	}

	Secret = os.Getenv("SECRET")
	if Secret == "" {
		log.Fatal("SECRET not set")
	}

	AdminKey = os.Getenv("ADMINKEY")
	if AdminKey == "" {
		log.Fatal("ADMINKEY not set")
	}
}
