package db

import (
	env "github.com/cesart18/ranking_app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitiDB() {

	var err error
	DB, err = gorm.Open(postgres.Open(env.DBUrl), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
