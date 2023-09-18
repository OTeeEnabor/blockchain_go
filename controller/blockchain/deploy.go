package blockchain

import (
	"encoding/json"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

// create a contract data struct
type ContractData struct {
	Data Data `json:"data`
}

// create Data struct
type Data struct {
	Bytecode Bytecode `json"bytecode`
}

// create contract Bytecode type
type Bytecode struct {
	Object string `json:"object"`
}

func DeployContract() (*hedera.TransactionRecord, error) {
	// connect to the hedera client using account id and private key from the development.yaml

	// initialise the client and check for err
	client, err := GetClient()

	if err != nil {
		return &hedera.TransactionRecord{},err 

	}

	//  read in the compiled contract from the contract.json file in contract/contract.json

	rawSmartContract, err := os.ReadFile("contract/contract.json") // will this work on server

	if err != nil {
		println(err.Error(),": error reading contract.json")
		return &hedera.TransactionRecord{}, err
	}

	// initialise the contract with its bytecode struct down used to extract the bytecode object

	var contract ContractData

	// Parse the bytecode object in contract.json 
	err = json.Unmarshal([]byte(rawSmartContract), &contract)

	//  check if there was an error in parsing the code
	if err != nil {
		println(err.Error(),":error unmarshaling")
		return &hedera.TransactionRecord{}, err
	}

	// Create the transaction
	// create file storing the bytecode and contract
	// set the amount of gas needed for the transaction
	// then part the bytecode object to it 
	contractCreate := hedera.NewContractCreateFlow().
	SetGas(100000).
	SetBytecode([]byte(contract.Data.Bytecode.Object))

	// sign the transaction with the hedera client key created
	// using the private key and submit to Hedera network
	txResponse, err := contractCreate.Execute(client)
	if err != nil {
		return &hedera.TransactionRecord{}, err
	}

	// get the transaction record of the contract deployed
	getRecord, err := txResponse.GetRecord(client)
	if err != nil{
		return &hedera.TransactionRecord{}, err
	}
	// access the value of getRecord using pointer notation
	record := &getRecord
	// return the record and nil since there will be no error at this point
	return record, nil
}