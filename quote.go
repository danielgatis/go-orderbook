package orderbook

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

var _ json.Marshaler = (*Quote)(nil)
var _ json.Unmarshaler = (*Quote)(nil)

// Quote represents a quote.
type Quote struct {
	price           decimal.Decimal
	remainingAmount decimal.Decimal
}

// NewQuote creates a new quote.
func NewQuote(price, remainingAmount decimal.Decimal) *Quote {
	return &Quote{price, remainingAmount}
}

// Price returns the price.
func (m *Quote) Price() decimal.Decimal {
	return m.price
}

// RemainingAmount returns the remaining amount.
func (m *Quote) RemainingAmount() decimal.Decimal {
	return m.remainingAmount
}

// MarshalJSON implements json.Marshaler.
func (m *Quote) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Price           decimal.Decimal `json:"price"`
			RemainingAmount decimal.Decimal `json:"remainingAmount"`
		}{
			m.price,
			m.remainingAmount,
		},
	)
}

// UnmarshalJSON implements json.Unmarshaler.
func (m *Quote) UnmarshalJSON(data []byte) error {
	obj := struct {
		Price           decimal.Decimal `json:"price"`
		RemainingAmount decimal.Decimal `json:"remainingAmount"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("Quote.Unmarshal(%s): %w", data, err)
	}

	m.price = obj.Price
	m.remainingAmount = obj.RemainingAmount

	return nil
}
