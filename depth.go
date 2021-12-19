package orderbook

import (
	"encoding/json"
	"fmt"
)

var _ json.Marshaler = (*Depth)(nil)
var _ json.Unmarshaler = (*Depth)(nil)

// Depth represents a order book depth.
type Depth struct {
	bids []*PriceLevel
	asks []*PriceLevel
}

// NewDepth creates a new depth.
func NewDepth(bids, asks []*PriceLevel) *Depth {
	return &Depth{bids, asks}
}

// Bids returns a range of price leves.
func (d *Depth) Bids() []*PriceLevel {
	return d.bids
}

// Asks returns a range of price leves.
func (d *Depth) Asks() []*PriceLevel {
	return d.asks
}

// MarshalJSON implements json.MarshalJSON.
func (d *Depth) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		&struct {
			Bids []*PriceLevel `json:"bids"`
			Asks []*PriceLevel `json:"asks"`
		}{
			d.bids,
			d.asks,
		},
	)
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *Depth) UnmarshalJSON(data []byte) error {
	obj := struct {
		Bids []*PriceLevel `json:"bids"`
		Asks []*PriceLevel `json:"asks"`
	}{}

	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("Depth.Unmarshal(%s): %w", data, err)
	}

	d.bids = obj.Bids
	d.asks = obj.Asks

	return nil
}
