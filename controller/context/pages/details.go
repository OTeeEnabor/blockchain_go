package pages

import (
	"fmt"
	"net/http"

	"github.com/OTeeEnabor/blockchain_go/controller/blockchain"
	"github.com/OTeeEnabor/blockchain_go/db"
	"github.com/OTeeEnabor/blockchain_go/models"
	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/labstack/echo/v4"
)

func DetailsContext(c echo.Context) error {
	// get the productId from the query parameter
	getProductId := c.Param("productId")
	fmt.Println(getProductId)
	// get db record with the product id
	// connect to the db
	db, err := db.ConnectToDb()
	// check if error
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	// initialise the product model
	var product models.Products
	// search the db for product id that equals to getProductId
	result := db.Where("product_id = ?", getProductId).First(&product)

	// if record is found, return it
	if result.Error == nil {
		// product found the record in the database
		// get contract id from the database
		// and unmarshal it using scan(automatically) in models/contract
		
		hederaContractID := hedera.ContractID(*product.ContractId)
		// use the unmarshal contract id and product id to query the smart contract in hedera
		// return the query result
		contractQuery, err := blockchain.GetContractRecord(hederaContractID, getProductId)

		// check if there is an error in query the smart contract
		if err != nil {
			panic(err) // stop the application
		}
		// pass the query result as context
		return c.Render(http.StatusOK, "detail.html", map[string]interface{}{
			"product_id": product.ProductID,
			"contract_id": product.ContractIdString,
			"gas_used": contractQuery.GasUsed,
			"transaction_fee": product.ChargeFee,
			"payer_account": product.PayerAccount,
		})
	}
	return err
}