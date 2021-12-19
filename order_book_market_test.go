package orderbook_test

import (
	"encoding/json"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/danielgatis/go-orderbook"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestProcessMarketOrderBids(t *testing.T) {
	type input struct {
		OrderID  string
		traderID string
		side     orderbook.Side
		amount   decimal.Decimal
		price    decimal.Decimal
	}

	type snapshot struct {
		Book   *orderbook.OrderBook
		Trades []*orderbook.Trade
		Err    string
	}

	tests := []struct {
		name  string
		input input
	}{}

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
							"amount": "5",
							"price": "500"
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
							"amount": "1",
							"price": "300"
						}
					]
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			trades, err := book.ProcessMarketOrder(tt.input.OrderID, tt.input.traderID, tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book:   &book,
				Trades: trades,
				Err:    errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestProcessMarketOrderAsks(t *testing.T) {
	type input struct {
		OrderID  string
		traderID string
		side     orderbook.Side
		amount   decimal.Decimal
		price    decimal.Decimal
	}

	type snapshot struct {
		Book   *orderbook.OrderBook
		Trades []*orderbook.Trade
		Err    string
	}

	tests := []struct {
		name  string
		input input
	}{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			given := []byte(`
				{
					"bids": [
						{
							"id": "1",
							"traderId": "1",
							"side": "buy",
							"amount": "5",
							"price": "500"
						},
						{
							"id": "2",
							"traderId": "2",
							"side": "buy",
							"amount": "1",
							"price": "400"
						},
						{
							"id": "3",
							"traderId": "3",
							"side": "buy",
							"amount": "0.5",
							"price": "300"
						}
					],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			trades, err := book.ProcessMarketOrder(tt.input.OrderID, tt.input.traderID, tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book:   &book,
				Trades: trades,
				Err:    errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}

func TestProcessMarketOrderValidations(t *testing.T) {
	type input struct {
		OrderID  string
		traderID string
		side     orderbook.Side
		amount   decimal.Decimal
		price    decimal.Decimal
	}

	type snapshot struct {
		Book   *orderbook.OrderBook
		Trades []*orderbook.Trade
		Err    string
	}

	tests := []struct {
		name  string
		input input
	}{
		{
			name: "invalid order id",
			input: input{
				OrderID:  "",
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(5),
				price:    decimal.NewFromInt(550),
			},
		},
		{
			name: "order already exists",
			input: input{
				OrderID:  "1",
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(5),
				price:    decimal.NewFromInt(550),
			},
		},
		{
			name: "invalid trader id",
			input: input{
				OrderID:  "4",
				traderID: "",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(5),
				price:    decimal.NewFromInt(550),
			},
		},
		{
			name: "invalid amount",
			input: input{
				OrderID:  "4",
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(0),
				price:    decimal.NewFromInt(550),
			},
		},
		{
			name: "invalid price",
			input: input{
				OrderID:  "4",
				traderID: "4",
				side:     orderbook.Sell,
				amount:   decimal.NewFromInt(5),
				price:    decimal.NewFromInt(0),
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
							"amount": "5",
							"price": "500"
						},
						{
							"id": "2",
							"traderId": "2",
							"side": "buy",
							"amount": "1",
							"price": "400"
						},
						{
							"id": "3",
							"traderId": "3",
							"side": "buy",
							"amount": "0.5",
							"price": "300"
						}
					],
					"asks": []
				}
			`)

			var book orderbook.OrderBook
			err := json.Unmarshal(given, &book)
			assert.Nil(t, err)

			trades, err := book.ProcessMarketOrder(tt.input.OrderID, tt.input.traderID, tt.input.side, tt.input.amount, tt.input.price)

			var errorStr string
			if err != nil {
				errorStr = err.Error()
			}

			s, err := json.Marshal(&snapshot{
				Book:   &book,
				Trades: trades,
				Err:    errorStr,
			})

			assert.Nil(t, err)
			cupaloy.SnapshotT(t, s)
		})
	}
}
