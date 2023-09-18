package blockchain

import (
	"github.com/OTeeEnabor/blockchain_go/config"
	"github.com/hashgraph/hedera-sdk-go/v2"
)

// set environment
var env = config.EnVar

// create function to set and get hedera client
func GetClient() (*hedera.Client, error) {
	// initialise the client and error to a variable
	// need to connect to the hedera network

	var client *hedera.Client
	var err error

	// Retrieve the network type from environment variables HEDERA_NETWORK
	// use the testnet

	client, err = hedera.ClientForName(env.GetString("HEDERA_NETWORK"))

	// check for successful connection is established
	if err != nil {
		println(err.Error(), ": error creating client")
		// return empty client and an error
		return &hedera.Client{}, err
	}

	//Attempt to retrieve account ID from environment variable in development.yaml
	operatorAccountID, err := hedera.AccountIDFromString(env.GetString("ACCOUNT_ID"))

	// check if there error in retrieving account from hedera  test network

	if err != nil {
		println(err.Error(), ": error to convert string to AccountID")
		// return an empty client and an error
		return &hedera.Client{},err
	}

	// retrieve private key from environment variable in development.yaml
	operatorKey, err := hedera.PrivateKeyFromString(env.GetString("DER_ENCODED_PRIVATE_KEY"))

	// check if there error in retrieving account from hedera test network
	if err != nil {

		println(err.Error(),": error to convert string to private Key")

		// return an empty client and error
		return &hedera.Client{}, err

	}

	// connect to the client using your accountID and private key
	client.SetOperator(operatorAccountID, operatorKey)

	return client, nil

}