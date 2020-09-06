package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bitmex_secret := os.Getenv("BITMEX_SECRET")
	bitmex_id := os.Getenv("BITMEX_ID")
	fmt.Println(bitmex_secret)
	fmt.Println(bitmex_id)
}
