package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BitmexID     string
	BitmexSecret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	BitmexID = os.Getenv("BITMEX_ID")
	BitmexSecret = os.Getenv("BITMEX_SECRET")
}

func main() {
	fmt.Println(BitmexID)
	fmt.Println(BitmexSecret)
}
