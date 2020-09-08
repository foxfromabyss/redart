package bitmex

// Message follows pattern {"op": "<command>", "args": ["arg1", "arg2", "arg3"]}
type Message struct {
	Op   string   `json:"op"`
	Args []string `json:"args"`
}

// SubscriptionResultMessage ...
type SubscriptionResultMessage struct {
	Success   bool    `json:"success"`
	Subscribe Pair    `json:"subscribe"`
	Request   Message `json:"request"`
}

// SubscriptionMessage ...
type SubscriptionMessage struct {
	Table  string `json:"table"`
	Action string `json:"action"` // enum [partial | insert | update | delete]
	Data   []struct {
		Symbol string  `json:"symbol"`
		ID     int     `json:"id"`
		Side   string  `json:"side"` // enum [Buy | Sell]
		Size   float32 `json:"size"`
		Price  float32 `json:"price"`
	} `json:"data"`
}

const (
	ActionPartial = "partial"
	ActionInsert  = "insert"
	ActionUpdate  = "update"
	ActionDelete  = "delete"

	SideBuy  = "Buy"
	SideSell = "Sell"
)

type SubscriptionMessageHandler func(SubscriptionMessage)
