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
	ServerPort = os.Getenv("SERVER_PORT")
	Secret = os.Getenv("SECRET")
	AdminKey = os.Getenv("ADMINKEY")
}
