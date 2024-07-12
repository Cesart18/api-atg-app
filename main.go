package main

import (
	"github.com/cesart18/ranking_app/config"
	"github.com/cesart18/ranking_app/db"
	"github.com/cesart18/ranking_app/routes"
)

func init() {
	config.InitEnvVariables()
	db.InitiDB()
}

func main() {
	r := routes.SetupRouter()

	r.Run()
}
