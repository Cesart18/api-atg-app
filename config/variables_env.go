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

func InitEnvVariables() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("error en la variable de entorno")
	}
	DBUrl = os.Getenv("DATABASE_URL")
	if DBUrl == "" {
		log.Fatal("DATABASE_URL not set")
	}

	ServerPort = os.Getenv("PORT")
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
