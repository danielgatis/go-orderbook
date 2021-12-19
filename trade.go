package orderbook

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

var _ json.Marshaler = (*Trade)(nil)
var _ json.Unmarshaler = (*Trade)(nil)

// Trade represents a match between a maker order and a taker order.
type Trade struct {
	takerOrderID string
	makerOrderID string
	amount       decimal.Decimal
	price        decimal.Decimal
}

// TakerOrderID returns the taker order id.
func (t *Trade) TakerOrderID() string {
	return t.takerOrderID
}

// MakerOrderID returns the maker order id.
func (t *Trade) MakerOrderID() string {
	return t.makerOrderID
}

// Amount returns the amount.
func (t *Trade) Amount() decimal.Decimal {
	return t.amount
}

// Price returns the price.
func (t *Trade) Price() decimal.Decimal {
	return t.price
}

// NewTrade creates a new trade.
func NewTrade(takerOrderID, makerOrderID string, amount, price decimal.Decimal) *Trade {
	return &Trade{takerOrderID, makerOrderID, amount, price}
}

// MarshalJSON implements json.Marshaler.
func (t *Trade) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			TakerOrderID string          `json:"takerOrderId"`
			MakerOrderID string          `json:"makerOrderId"`
			Amount       decimal.Decimal `json:"amount"`
			Price        decimal.Decimal `json:"price"`
		}{
			t.takerOrderID,
			t.makerOrderID,
			t.amount,
			t.price,
		},
	)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Trade) UnmarshalJSON(data []byte) error {
	obj := struct {
		TakerOrderID string          `json:"takerOrderId"`
		MakerOrderID string          `json:"makerOrderId"`
		Amount       decimal.Decimal `json:"amount"`
		Price        decimal.Decimal `json:"price"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("Trade.Unmarshal(%s): %w", data, err)
	}

	t.takerOrderID = obj.TakerOrderID
	t.makerOrderID = obj.MakerOrderID
	t.amount = obj.Amount
	t.price = obj.Price

	return nil
}
