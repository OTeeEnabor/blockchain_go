package db

import (
	"github.com/OTeeEnabor/blockchain_go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init env variable
var env = config.EnVar

func ConnectToDb()(*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(env.GetString("DNS")), &gorm.Config{})

	if err != nil {
		panic("fail to connect to database")
	}
	return db, err
}