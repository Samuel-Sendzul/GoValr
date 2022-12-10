package main

import (
	//"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/govalr/account"
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


	account.GetBalances(os.Getenv("key"), os.Getenv("secret"))
	// key := os.Getenv("key")
	// secret := os.Getenv("secret")



}