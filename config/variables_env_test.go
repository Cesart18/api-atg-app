package config_test

import (
	"os"
	"testing"

	"github.com/cesart18/ranking_app/config"
	"github.com/stretchr/testify/assert"
)

func TestInitEnvVariables(t *testing.T) {

	config.InitEnvVariables("../.env")

	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")
	ServerPort := os.Getenv("SERVER_PORT")

	assert.Equal(t, DBHost, config.DBHost)
	assert.Equal(t, DBUser, config.DBUser)
	assert.Equal(t, DBPassword, config.DBPassword)
	assert.Equal(t, DBName, config.DBName)
	assert.Equal(t, DBPort, config.DBPort)
	assert.Equal(t, ServerPort, config.ServerPort)

}
