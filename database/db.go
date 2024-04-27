package database

import (
	"log"
	"playground/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	if db, err = gorm.Open(postgres.Open(util.Config.DSN), &gorm.Config{}); err != nil {
		log.Fatal("unable to connect to db")
	}
}
