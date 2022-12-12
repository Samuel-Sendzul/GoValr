package private

import (
	"encoding/json"
	"fmt"
	"log"
)

type DepositAddress struct {
	Currency string `json:"currency"`
	Address string `json:"address"`
	Tag string `json:"paymentReference"`
}

type Deposit struct {
	CurrencyCode string `json:"currencyCode"`
	ReceiveAddress string `json:"receiveAddress"`
	TransactionHash string `json:"transactionHash"`
	Amount string `json:"amount"`
	CreatedAt string `json:"createdAt"`
	Confirmations int `json:"confirmations"`
	Confirmed bool `json:"confirmed"`
	ConfirmedAt string `json:"confirmedAt"`
}

// GetDepositAddress gets the deposit address and tag (if applicable) for a specified
// currency code.
func GetDepositAddress(key string, secret string, currencyCode string) (DepositAddress, error) {
	resp, err := SignedGetRequest(key, secret, fmt.Sprintf("/v1/wallet/crypto/%s/deposit/address", currencyCode), nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error perfroming the signed GET request.: %s", err.Error()))
	}

	defer resp.Body.Close()

	var address DepositAddress
	json.NewDecoder(resp.Body).Decode(&address)

	return address, nil
}


// GetDepositHistory gets the deposit history for a given currency code.
func GetDepositHistory(key string, secret string, currencyCode string) ([]Deposit, error){
	resp, err := SignedGetRequest(key, secret,  fmt.Sprintf("/v1/wallet/crypto/%s/deposit/history", currencyCode), nil)
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error perfroming the signed GET request.: %s", err.Error()))
	}

	defer resp.Body.Close()

	var depositHistory []Deposit
	json.NewDecoder(resp.Body).Decode(&depositHistory)

	return depositHistory, nil
}
