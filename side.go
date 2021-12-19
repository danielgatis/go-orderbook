package orderbook

import (
	"encoding/json"
	"reflect"
)

var _ json.Marshaler = (*Side)(nil)
var _ json.Unmarshaler = (*Side)(nil)

// A Side of the order.
type Side int

const (
	// Sell for asks
	Sell Side = 0

	// Buy for bids
	Buy Side = 1
)

// String implements fmt.Stringer.
func (s Side) String() string {
	if s == Buy {
		return "buy"
	}

	return "sell"
}

// MarshalJSON implements json.Marshaler.
func (s Side) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *Side) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"buy"`:
		*s = Buy
	case `"sell"`:
		*s = Sell
	default:
		return &json.UnsupportedValueError{
			Value: reflect.New(reflect.TypeOf(data)),
			Str:   string(data),
		}
	}

	return nil
}
