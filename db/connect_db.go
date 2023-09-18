package db

import (
	"github.com/OTeeEnabor/blockchain_go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initialise env variable
var env = config.EnVar

//  create function to connect to database
func ConnectToDb()(*gorm.DB, error){
	// open the a connection to the database
	// this method creates a db connection and error if connection not successful
	db, err := gorm.Open(postgres.Open(env.GetString("DNS")), &gorm.Config{})
	// if there is an error
	if err != nil {
		// stop the program
		panic("fail to connect to database")
	}
	// if no error, return db, error (which is nil in this case)
	return db, err
}