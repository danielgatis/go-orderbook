package orderbook_test

import (
	"encoding/json"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/danielgatis/go-orderbook"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestQuoteBids(t *testing.T) {
	type input struct {
		traderID string
		side     orderbook.Side
		amount   decimal.Decimal
	}

	type snapshot struct {
		Book   *orderbook.OrderBook
		Result *orderbook.Quote
		Err    string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "happy path - 1",
			input: input{
				traderID: "4",
				side:     orderbook.Buy,
				amount:   decimal.NewFromInt(6),
			},
		},
		{
			name: "happy path - 2",
			input: input{
				traderID: "4",
				side:     orderbook.Buy,
				amount:   decimal.NewFromInt(1),
			},
		},
		{
			name: "empty book",
			input: input{
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(6),
			},
		},
		{
			name: "skip same trader",
			input: input{
				traderID: "2",
				side:     orderbook.Buy,
				amount:   decimal.NewFromInt(6),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [],
					"asks": [
						{
							"id": "1",
							"traderId": "1",
							"side": "sell",
							"amount": "2",
							"price": "200"
						},
						{
							"id": "2",
							"traderId": "2",
							"side": "sell",
							"amount": "2",
							"price": "400"
						},
						{
							"id": "3",
							"traderId": "3",
							"side": "sell",
							"amount": "2",
							"price": "600"
						}
					]
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			result, err := book.Quote(tt.input.traderID, tt.input.side, tt.input.amount)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book:   &book,
				Result: result,
				Err:    errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestQuoteAsks(t *testing.T) {
	type input struct {
		traderID string
		side     orderbook.Side
		amount   decimal.Decimal
	}

	type snapshot struct {
		Book   *orderbook.OrderBook
		Result *orderbook.Quote
		Err    string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "happy path - 1",
			input: input{
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(6),
			},
		},
		{
			name: "happy path - 2",
			input: input{
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(1),
			},
		},
		{
			name: "empty book",
			input: input{
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(6),
			},
		},
		{
			name: "skip same trader",
			input: input{
				traderID: "2",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(6),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [
						{
							"id": "1",
							"traderId": "1",
							"side": "buy",
							"amount": "2",
							"price": "200"
						},
						{
							"id": "2",
							"traderId": "2",
							"side": "buy",
							"amount": "2",
							"price": "400"
						},
						{
							"id": "3",
							"traderId": "3",
							"side": "buy",
							"amount": "2",
							"price": "600"
						}
					],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			result, err := book.Quote(tt.input.traderID, tt.input.side, tt.input.amount)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book:   &book,
				Result: result,
				Err:    errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestQuoteValidations(t *testing.T) {
	type input struct {
		traderID string
		side     orderbook.Side
		amount   decimal.Decimal
	}

	type snapshot struct {
		Book   *orderbook.OrderBook
		Result *orderbook.Quote
		Err    string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "invalid trader id",
			input: input{
				traderID: "",
				side:     orderbook.Buy,
				amount:   decimal.NewFromInt(6),
			},
		},
		{
			name: "invalid amount",
			input: input{
				traderID: "4",
				side:     orderbook.Buy,
				amount:   decimal.Zero,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [],
					"asks": [
						{
							"id": "1",
							"traderId": "1",
							"side": "sell",
							"amount": "2",
							"price": "200"
						},
						{
							"id": "2",
							"traderId": "2",
							"side": "sell",
							"amount": "2",
							"price": "400"
						},
						{
							"id": "3",
							"traderId": "3",
							"side": "sell",
							"amount": "2",
							"price": "600"
						}
					]
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			result, err := book.Quote(tt.input.traderID, tt.input.side, tt.input.amount)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book:   &book,
				Result: result,
				Err:    errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}
