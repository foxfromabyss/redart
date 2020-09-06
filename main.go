package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	BitmexID     string
	BitmexSecret string
	Address      string
)

// follows pattern {"op": "<command>", "args": ["arg1", "arg2", "arg3"]}
type BitmexMessage struct {
	Op   string
	Args []string
}

func bitmexMessageSubscription(subName string) BitmexMessage {
	args := []string{subName}
	message := BitmexMessage{Op: "subscribe", Args: args}
	return message
}

// func (b *BitmexMessage) json() string {
// 	result := ""
// 	result = result + "{\"op\": \"" + b.op + "\", "
// 	return math.Pi * c.r * c.r
// }

var addr = flag.String("addr", "testnet.bitmex.com", "http service address")

func listen() {
	addSubscribe := make(chan BitmexMessage)
	message := bitmexMessageSubscription("orderBookL2_25:XBTUSD")
	addSubscribe <- message

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{Scheme: "wss", Host: Address, Path: "/realtime"}

	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()
	for {
		select {
		// case <-done:
		// 	return
		case msg := <-addSubscribe:
			data, _ := json.Marshal(msg)
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, string(data)))
			if err != nil {
				log.Println("write close:", err)
				return
			}

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	BitmexID = os.Getenv("BITMEX_ID")
	BitmexSecret = os.Getenv("BITMEX_SECRET")
	Address = os.Getenv("BITMEX_ADDRESS")
}

func main() {
	fmt.Println(BitmexID)
	fmt.Println(BitmexSecret)

	listen()
}
