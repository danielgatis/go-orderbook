package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/danielgatis/go-orderbook"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func main() {
	book := orderbook.NewOrderBook("BTC/BRL")

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		side := []orderbook.Side{orderbook.Buy, orderbook.Sell}[rand.Intn(2)]

		book.ProcessPostOnlyOrder(uuid.New().String(), uuid.New().String(), side, decimal.NewFromInt(rand.Int63n(1000)), decimal.NewFromInt(rand.Int63n(1000)))
	}

	depth, _ := json.Marshal(book.Depth())
	var buf bytes.Buffer
	json.Indent(&buf, depth, "", "  ")
	fmt.Println(buf.String())
}
