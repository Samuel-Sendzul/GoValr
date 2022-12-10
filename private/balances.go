package private

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"time"
)

var baseUrl = "https://api.valr.com/"

type Balance struct {
	Currency string 
}

func GetBalances(key string, secret string) ([]Balance) {

	requestPath := fmt.Sprintf("%sv1/account/balances", baseUrl)

	SignRequest(secret, time.Now(), "GET", requestPath, "")


	resp, err := http.Get(requestPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Request error: %s", err))
	}

	defer resp.Body.Close()

	var balances []Balance
	json.NewDecoder(resp.Body).Decode(&balances)

	return balances
}
