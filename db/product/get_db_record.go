package product

import (
	"fmt"

	database "github.com/OTeeEnabor/blockchain_go/db"
	"github.com/OTeeEnabor/blockchain_go/models"
)

func GetProductDetails(productID string) (map[string]interface{}, error) {

	db, err := database.ConnectToDb()

	if err != nil {
		fmt.Println("Failed to connect to database")
		return map[string]interface{}{}, err
	}

	var product models.Products
	// search product
	result := db.Where("product_id = ?", productID).First(&product);

	if result.Error == nil {
		return map[string]interface{}{
			"product_id":         product.ProductID,
			"colour":             product.Colour,
			"quantity":           product.Quantity,
			"timestamp":          product.Timestamp,
			"contract_id":        product.ContractIdString,
			"gas_used":           product.GasUsed,
			"transaction_id":     product.TransactionId,
			"transaction_fee":    product.ChargeFee,
			"payer_account":      product.PayerAccount,
			"transaction_status": product.Status,
		}, nil	
	}
	return map[string]interface{}{}, err
}