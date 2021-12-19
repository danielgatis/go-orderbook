package orderbook

import (
	"container/list"
	"encoding/json"
	"fmt"
	"sync"
)

var _ json.Marshaler = (*OrderBook)(nil)
var _ json.Unmarshaler = (*OrderBook)(nil)

// OrderBook represents a order book for a given market symbol.
type OrderBook struct {
	sync.RWMutex
	symbol  string
	version uint64
	orders  map[string]*list.Element
	asks    *OrderSide
	bids    *OrderSide
}

// NewOrderBook creates a new order book.
func NewOrderBook(symbol string) *OrderBook {
	return &OrderBook{sync.RWMutex{}, symbol, 0, make(map[string]*list.Element), NewOrderSide(Sell), NewOrderSide(Buy)}
}

// Symbol returns the symbol.
func (ob *OrderBook) Symbol() string {
	return ob.symbol
}

// Version returns the version. The version is auto incremented by each change.
func (ob *OrderBook) Version() uint64 {
	return ob.version
}

// Reset resets the order book.
func (ob *OrderBook) Reset() {
	defer ob.Unlock()
	ob.Lock()

	ob.orders = make(map[string]*list.Element)
	ob.asks = NewOrderSide(Sell)
	ob.bids = NewOrderSide(Buy)
	ob.version = 0
}

// MarshalJSON implements json.MarshalJSON.
func (ob *OrderBook) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Symbol  string   `json:"symbol"`
			Bids    []*Order `json:"bids"`
			Asks    []*Order `json:"asks"`
			Version uint64   `json:"version"`
		}{
			ob.symbol,
			ob.bids.Orders(),
			ob.asks.Orders(),
			ob.version,
		},
	)
}

// UnmarshalJSON implements json.Unmarshaler.
func (ob *OrderBook) UnmarshalJSON(data []byte) error {
	defer ob.Unlock()
	ob.Lock()

	obj := struct {
		Symbol  string   `json:"symbol"`
		Bids    []*Order `json:"bids"`
		Asks    []*Order `json:"asks"`
		Version uint64   `json:"version"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("OrderBook.Unmarshal(%s): %w", data, err)
	}

	ob.symbol = obj.Symbol
	ob.version = obj.Version
	ob.orders = make(map[string]*list.Element)

	ob.asks = NewOrderSide(Sell)
	for _, order := range obj.Asks {
		ob.orders[order.id] = ob.asks.Append(order)
	}

	ob.bids = NewOrderSide(Buy)
	for _, order := range obj.Bids {
		ob.orders[order.id] = ob.bids.Append(order)
	}

	return nil
}
