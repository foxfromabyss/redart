package main

import (
	"fmt"
	"github.com/foxfromabyss/redart/bitmex"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	BitmexID      string
	BitmexSecret  string
	BitmexAddress string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	BitmexID = os.Getenv("BITMEX_ID")
	BitmexSecret = os.Getenv("BITMEX_SECRET")
	BitmexAddress = os.Getenv("BITMEX_ADDRESS")
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	bitmex.Connect(BitmexID, BitmexSecret, BitmexAddress)
	bitmex.AddSubscriptionMessageHandler(func(message bitmex.SubscriptionMessage) {
		fmt.Println(message)
	})
	bitmex.Subscribe(bitmex.PairL225XBTUSD)

waitLoop:
	for {
		select {
		case err := <-bitmex.Notify():
			log.Fatal(err)
		case sig := <-signals:
			log.Printf("application stopped (system signal %s)", sig.String())
			break waitLoop
		}
	}

	bitmex.Close()
}
