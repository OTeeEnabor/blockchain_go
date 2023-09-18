package blockchain

import (
	// "github.com/hashgraph/hedera-sdk-go"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

// query the get_record function

func GetContractRecord(newContractID hedera.ContractID, productID string) (hedera.ContractFunctionResult, error){

	// initialise the client and connect to it
	client, err := GetClient()

	// check if there is an error connecting to the client
	if err != nil {
		return hedera.ContractFunctionResult{}, err
	}

	// close the connection at end of function
	// defer keyword is used to do this 
	defer func() {
		err = client.Close()

		// check if there is an error in closing connection to te client
		if err != nil {
			println(err.Error(),":error closing client")
			return
		}
	}()
	// set parameters for contract from solidity contract/contract.sol
	// get_recird requires the product ID parameter
	// here, just init the parameter, yet to attach the parameter to the function call

	contractFunctionParams := hedera.NewContractFunctionParameters().AddString(productID)

	// call a method on a contract that exists on network, and attach the parameter to it
	contractQuery, err := hedera.NewContractCallQuery().
	// set which contract id to use from the deployed contract
	SetContractID(newContractID).
	// set the gas to use
	SetGas(100000).
	// set the query payment explicitly since sometimes automatic payment
	// calculated may be too low for the transaction
	SetQueryPayment(hedera.NewHbar(1)).
	// set the function to call on the contract
	// along which the parameter required
	SetFunction("get_record", contractFunctionParams).
	Execute(client)

	// check if the function call returned an error
	if err != nil {
		println(err.Error(),": error executing the contract call query")
		return hedera.ContractFunctionResult{},err
	}

	// return the smart contract query
	return contractQuery, err
	
}
