package orderbook

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Restore restores a new order book from raw representation. Raw format is a nested arrays like: raw[ask[price][amount]][bid[price][amount]][symbol].
func Restore(version uint64, raw [][][]string) *OrderBook {
	book := NewOrderBook(raw[0][0][1])
	book.version = version

	for _, ask := range raw[1] {
		id := uuid.New()
		price := decimal.RequireFromString(ask[0])
		amount := decimal.RequireFromString(ask[1])
		book.ProcessPostOnlyOrder(id.String(), id.String(), Sell, amount, price)
	}

	for _, bid := range raw[2] {
		id := uuid.New()
		price := decimal.RequireFromString(bid[0])
		amount := decimal.RequireFromString(bid[1])
		book.ProcessPostOnlyOrder(id.String(), id.String(), Buy, amount, price)
	}

	return book
}
