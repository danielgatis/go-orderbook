package orderbook_test

import (
	"encoding/json"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/danielgatis/go-orderbook"
	"github.com/stretchr/testify/assert"
)

func TestDethBids(t *testing.T) {
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

	result := book.Depth()
	s, err := json.Marshal(result)

	assert.Nil(t, err)
	cupaloy.SnapshotT(t, s)
}

func TestDepthAsks(t *testing.T) {
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

	result := book.Depth()
	s, err := json.Marshal(result)

	assert.Nil(t, err)
	cupaloy.SnapshotT(t, s)
}
