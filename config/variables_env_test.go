package config_test

import (
	"os"
	"testing"

	"github.com/cesart18/ranking_app/config"
	"github.com/stretchr/testify/assert"
)

func TestInitEnvVariables(t *testing.T) {

	config.InitEnvVariables("../.env")

	DBUrl := os.Getenv("DATABASE_URL")
	ServerPort := os.Getenv("SERVER_PORT")

	assert.Equal(t, DBUrl, config.DBUrl)
	assert.Equal(t, ServerPort, config.ServerPort)

}
