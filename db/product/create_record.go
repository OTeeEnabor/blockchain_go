package product

import (
	"fmt"
	"time"

	"github.com/OTeeEnabor/blockchain_go/db"
	"github.com/OTeeEnabor/blockchain_go/models"
)

func CreateRecord(productID string , cornColour string, quantity int64, timeStamp time.Time) error{

	// connect to database
	db, err := db.ConnectToDb()
	
	// if error occurred 
	if err != nil {
		// print out error message
		fmt.Println("Failed to connect to database")
		// return err
		return err
	}

	// migrate the schema, 
	// create table - if table is not in database
	db.AutoMigrate(&models.Products{})

	// Create the record and store it in the Products table
	db.Create(&models.Products{
		ProductID: productID,
		Colour: cornColour,
		Quantity: quantity,
		Timestamp: timeStamp,
	})

	return nil

}