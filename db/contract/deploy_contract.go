package contract

import (
	"fmt"
	"time"

	"github.com/OTeeEnabor/blockchain_go/controller/blockchain"
	database "github.com/OTeeEnabor/blockchain_go/db"
	model "github.com/OTeeEnabor/blockchain_go/models"
)

// save deploy contract details in db
func DeployContract(productID string) model.Contract {
	// connect to the database
	db, err := database.ConnectToDb()

	// check if connection fails 
	if err != nil {
		panic(err)
	}

	// initialise the contract model
	var modelContract model.Contract

	// check if a contract has been deployed
	dbRecord := db.First(&modelContract)

	if dbRecord.Error != nil {
		// no deployment found

		// migrate the schema to create the contract table
		db.AutoMigrate(&model.Contract{})

		// deploy the contract
		contract, _ := blockchain.DeployContract()

		// unmarshal contract id record
		// contract id from hedera is in pointer format
		// unmarshal to use it

		cID := model.ContractID(*contract.Receipt.ContractID)

		// create contract db record using form
		db.Create(&model.Contract{
			Id:            productID,
			ContractId:    &cID,
			GasUsed:       contract.CallResult.GasUsed,
			TransactionId: fmt.Sprint(contract.TransactionID),
			Timestamp:     contract.ConsensusTimestamp,
			ChargeFee:     fmt.Sprint(contract.TransactionFee),
			PayerAccount:  fmt.Sprint(contract.TransactionID.AccountID),
			Status:        fmt.Sprint(contract.Receipt.Status),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		})
	}else {
		// call search query for the first record
		db.First(&modelContract)
	}

	return modelContract


}