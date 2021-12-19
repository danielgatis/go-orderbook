package orderbook

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

var _ json.Marshaler = (*PriceLevel)(nil)
var _ json.Unmarshaler = (*PriceLevel)(nil)

// PriceLevel takes a count of how many assets have that price.
type PriceLevel struct {
	price  decimal.Decimal
	amount decimal.Decimal
}

// NewPriceLevel creates a new price level.
func NewPriceLevel(price, amount decimal.Decimal) *PriceLevel {
	return &PriceLevel{price, amount}
}

// Price returns the price.
func (p *PriceLevel) Price() decimal.Decimal {
	return p.price
}

// Amount returns the amount.
func (p *PriceLevel) Amount() decimal.Decimal {
	return p.amount
}

// MarshalJSON implements json.Marshaler.
func (p *PriceLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Amount decimal.Decimal `json:"amount"`
			Price  decimal.Decimal `json:"price"`
		}{
			p.amount,
			p.price,
		},
	)
}

// UnmarshalJSON implements json.Unmarshaler.
func (p *PriceLevel) UnmarshalJSON(data []byte) error {
	obj := struct {
		Amount decimal.Decimal `json:"amount"`
		Price  decimal.Decimal `json:"price"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("PriceLevel.Unmarshal(%s): %w", data, err)
	}

	p.amount = obj.Amount
	p.price = obj.Price

	return nil
}
