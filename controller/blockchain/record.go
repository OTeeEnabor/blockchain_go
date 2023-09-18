package blockchain

import "github.com/hashgraph/hedera-sdk-go/v2"

func SetContractRecord(newContractID hedera.ContractID, productID string, colour string, quantity int64) (*hedera.TransactionRecord, error) {

	//  initialise the client and error variables
	client, err := GetClient()

	if err != nil {
		return &hedera.TransactionRecord{},err
	}
	// close the client after running with the defer keyword
	// this closes the connection when the function execution ends
	// inefficient to keep client connected
	defer func() {
		err =  client.Close()
		if err != nil {
			println(err.Error(),": error closing the client")
			return
		}
	}()
	// set parameters for the contract from solidity contract/contract.sol
	// set record requires the productid parameter
	// here initialise the parameter but yet to attach it to the function call
	contractFunctionParams := hedera.NewContractFunctionParameters().
	AddString(productID).
	AddString(colour).
	AddInt64(quantity)

	// add the record
	contractExecuteID, err := hedera.NewContractExecuteTransaction().
	// set contract id deploy contract from
	SetContractID(newContractID).
	// set the gas fee to execute the contract call
	SetGas(2000000).
	// Set the function to call the parameters to send 
	SetFunction("set_record",contractFunctionParams).
	Execute(client)

	// check if the function call returned an error
	if err !=nil {
		println(err.Error(),":error executing the contract")
		return &hedera.TransactionRecord{},err
	}

	// get the record to make sure the transaction was executed
	contractExecuteRecord, err := contractExecuteID.GetRecord(client)

	// check if there is an error 
	if err != nil{
		println(err.Error(),":error retrieving contract execution record")
		return &hedera.TransactionRecord{},err

	}
	// get transaction record
	record := &contractExecuteRecord

	return record, nil
}