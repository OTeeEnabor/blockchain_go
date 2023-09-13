package pages

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	database "github.com/OTeeEnabor/blockchain_go/db/product"
	"github.com/labstack/echo/v4"
)

func FormContext(c echo.Context) error {
	// get productID form entry
	productID := c.FormValue("product_id")

	// get quantity form entry
	quantity := c.FormValue("quantity")

	//  get colour form entry
	cornColour :=c.FormValue("corn_colour")
	fmt.Printf("%v,%v,%v", productID, quantity,cornColour)

	//  convert quantity from string to correct int64
	numQuantity, _ := strconv.ParseInt(quantity, 10, 64)

	// set to current time
	currentTime := time.Now()

	// create a record in the database with the form values and current time
	database.CreateRecord(productID, cornColour, numQuantity, currentTime)
	
	// print record saved
	path := fmt.Sprintf("/%v",productID)

	return c.Redirect(http.StatusMovedPermanently, path)
}