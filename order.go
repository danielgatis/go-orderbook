package orderbook

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

var _ json.Marshaler = (*Order)(nil)
var _ json.Unmarshaler = (*Order)(nil)

// Order represents a bid or an ask.
type Order struct {
	id       string
	traderID string
	side     Side
	amount   decimal.Decimal
	price    decimal.Decimal
}

// NewOrder creates a new order.
func NewOrder(ID, traderID string, side Side, amount, price decimal.Decimal) *Order {
	return &Order{ID, traderID, side, amount, price}
}

// ID returns the order ID.
func (o *Order) ID() string {
	return o.id
}

// TraderID returns the trader ID.
func (o *Order) TraderID() string {
	return o.traderID
}

// Side returns the side.
func (o *Order) Side() Side {
	return o.side
}

// Amount returns the amount.
func (o *Order) Amount() decimal.Decimal {
	return o.amount
}

// Price returns the price.
func (o *Order) Price() decimal.Decimal {
	return o.price
}

// MarshalJSON implements json.Marshaler.
func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			ID       string          `json:"id"`
			TraderID string          `json:"traderId"`
			Side     Side            `json:"side"`
			Amount   decimal.Decimal `json:"amount"`
			Price    decimal.Decimal `json:"price"`
		}{
			o.id,
			o.traderID,
			o.side,
			o.amount,
			o.price,
		},
	)
}

// UnmarshalJSON implements json.Unmarshaler.
func (o *Order) UnmarshalJSON(data []byte) error {
	obj := struct {
		ID       string          `json:"id"`
		TraderID string          `json:"traderId"`
		Side     Side            `json:"side"`
		Amount   decimal.Decimal `json:"amount"`
		Price    decimal.Decimal `json:"price"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("Order.Unmarshal(%s): %w", data, err)
	}

	o.id = obj.ID
	o.side = obj.Side
	o.traderID = obj.TraderID
	o.side = obj.Side
	o.amount = obj.Amount
	o.price = obj.Price

	return nil
}
