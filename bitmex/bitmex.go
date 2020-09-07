package bitmex

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var (
	connection   *websocket.Conn
	addSubscribe chan Message
	errorsChan   chan error

	handlers []SubscriptionMessageHandler
)

// Connect connects to bitmex websocket API
func Connect(ID string, Secret string, Address string) {
	addSubscribe = make(chan Message)
	errorsChan = make(chan error)

	handlers = make([]SubscriptionMessageHandler, 0)

	bitmexURL := url.URL{Scheme: "wss", Host: Address, Path: "/realtime"}
	log.Printf("bitmex connecting to %s", bitmexURL.String())
	var err error
	connection, _, err = websocket.DefaultDialer.Dial(bitmexURL.String(), nil)
	if err != nil {
		errorsChan <- err
	}

	// TODO: stop goroutines properly (on interrupt)
	go func() {
		for {
			_, data, err := connection.ReadMessage()
			if err != nil {
				errorsChan <- err
				log.Println("read:", err)
				return
			}

			var message SubscriptionMessage
			err = json.Unmarshal(data, &message)
			if err != nil {
				errorsChan <- err
				log.Println("unmarshal:", err)
			}

			for _, handler := range handlers {
				handler(message)
			}
		}
	}()

	go func() {
		for {
			select {
			case message := <-addSubscribe:
				data, _ := json.Marshal(message)
				err := connection.WriteMessage(websocket.TextMessage, data)
				if err != nil {
					errorsChan <- err
				}
			}
		}
	}()
}

// Close closes connection to bitmex websocket API
func Close() {
	log.Println("bitmex closing connection")
	err := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		errorsChan <- err
		return
	}

	err = connection.Close()
	if err != nil {
		errorsChan <- err
	}
}

// Notify returns a channel to notify caller about errors
func Notify() <-chan error {
	return errorsChan
}

// Subscribe subscribes to specified pair
func Subscribe(pair Pair) {
	message := Message{
		Op:   "subscribe",
		Args: []string{string(pair)},
	}
	addSubscribe <- message
}

// AddSubscriptionMessageHandler adds handler which will be called back on each SubscriptionMessage
func AddSubscriptionMessageHandler(handler SubscriptionMessageHandler) {
	handlers = append(handlers, handler)
}
