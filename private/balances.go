package private

import (
"encoding/json"
	"log"
)

type Balance struct {
	Currency string `json:"currency"`
	Available string `json:"available"`
	Reserved string `json:"reserved"`
	Total string `json:"total"`
	UpdatedAt string `json:"updatedAt"`
	LendReserved string `json:"lendReserved"`
	BorrowReserved string `json:"borrowReserved"`
	BorrowedAmount string `json:"borrowedAmount"` 
}

func GetBalances(key string, secret string) ([]Balance, error) {

	resp, err := SignedGetRequest(key, secret, "/v1/account/balances", nil)
	if err != nil {
		log.Fatalf("Error perfroming the signed GET request.")
	}


	var balances []Balance
	json.NewDecoder(resp.Body).Decode(&balances)

	return balances, nil
}
