package pages

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/OTeeEnabor/blockchain_go/controller/blockchain"
	"github.com/OTeeEnabor/blockchain_go/db/contract"
	"github.com/OTeeEnabor/blockchain_go/db/product"
	"github.com/hashgraph/hedera-sdk-go/v2"
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

	// read contract
	// deploy contract if it has not been deployed before
	record := contract.DeployContract(productID)
	hederaContractID := hedera.ContractID(*record.ContractId)

	// save record in hedera blockchain
	blockchainRecord, err := blockchain.SetContractRecord(hederaContractID, productID, cornColour, numQuantity)

	// check for error in saving the blockchain record
	if err != nil{
		panic(err)
	}

	// get contract id of transaction
	contractId := blockchainRecord.Receipt.ContractID // returns a pointer

	// get string format of contract id
	contractIdStr := fmt.Sprint(blockchainRecord.Receipt.ContractID)

	// get the gas used for the transaction 
	gasUsed := blockchainRecord.CallResult.GasUsed

	// get transaction id for the transaction
	transactionId := fmt.Sprint(blockchainRecord.TransactionID)

	// transaction timestamp
	timestamp := blockchainRecord.ConsensusTimestamp

	// transaction fee 
	chargeFee := fmt.Sprint(blockchainRecord.TransactionFee)

	// fee payer account
	payerAccount := fmt.Sprint(blockchainRecord.TransactionID.AccountID)

	//status of transaction
	status := fmt.Sprint(blockchainRecord.Receipt.Status)

	//save product record in db
	err = product.CreateRecord(productID, cornColour, numQuantity, timestamp, contractId, contractIdStr, gasUsed, transactionId, chargeFee, payerAccount, status)

	//check if there is error saving in the database
	if err != nil {
		panic(err)
	}

	//set path
	path := fmt.Sprintf("/%v", productID)

	//direct to details page
	return c.Redirect(http.StatusMovedPermanently, path)
}