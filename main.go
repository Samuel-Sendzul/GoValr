package main

import (
	//"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	valr "github.com/govalr/private"
)


func loadDotEnv() {
	err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }
}

func main() {
	loadDotEnv()
	log.SetPrefix("govalr: ")


	valr.GetDepositHistory(os.Getenv("key"), os.Getenv("secret"), "ETH")
	//valr.GetBalances(os.Getenv("key"), os.Getenv("secret"))

	// key := os.Getenv("key")
	// secret := os.Getenv("secret")



}