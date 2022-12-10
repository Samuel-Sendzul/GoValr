package public

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
)

type OrderBookSide struct {
	Side string `json:"side"`
	Quantity string `json:"quantity"`
	Price string `json:"price"`
	CurrencyPair string `json:"currencyPair"`
	OrderCount int `json:"orderCount"`
}

type OrderBook struct {
	Asks []OrderBookSide `json:"Asks"`
	Bids []OrderBookSide `json:"Bids"`
	LastChange string `json:"LastChange"`
	SequenceNumber int  `json:"SequenceNumber"`

}

type Currency struct {

}

var baseUrl = "https://api.valr.com/"

func GetOrderBook(pair string) (OrderBook) {
	resp, err := http.Get(fmt.Sprintf("%sv1/public/%s/orderbook", baseUrl, pair))
	if err != nil {
		log.Fatal(fmt.Sprintf("Request error: %s", err))
	}

	defer resp.Body.Close()

	var orderbook OrderBook
	json.NewDecoder(resp.Body).Decode(&orderbook)

	return orderbook
}


