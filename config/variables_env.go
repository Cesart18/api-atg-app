package config

import (
	"fmt"
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

	fmt.Println("DATABASE_URL:", DBUrl)
	fmt.Println("SERVER_PORT:", ServerPort)
	fmt.Println("SECRET:", Secret)
	fmt.Println("ADMINKEY:", AdminKey)

}
