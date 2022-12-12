package private

import (
	"strconv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var baseUrl = "https://api.valr.com/"

func SignedGetRequest(key string, secret string, path string, params map[string]string) (*http.Response, error) {
	now := time.Now()
	nowTimesetamp := strconv.FormatInt(now.UnixNano()/1000000, 10)
	
	sig := SignRequest(secret, now, "GET", path, "")


	client := http.Client{}

	// Create request
	request, err := http.NewRequest("GET", fmt.Sprintf("%sv1/account/balances", baseUrl), nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Request error: %s", err))
	}
	request.Header.Set("X-VALR-API-KEY", key)
	request.Header.Set("X-VALR-SIGNATURE", sig)
	request.Header.Set("X-VALR-TIMESTAMP",  nowTimesetamp)


	// Send requests
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(fmt.Sprintf("Response error: %s", err))
	}

	// Handle unaurthorised requesting
	if resp.Status == "401 Unauthorized" {
		return nil, errors.New("401 Unauthorized")
	}

	return resp, nil
}

