package orderbook

import (
	"strings"

	"github.com/shopspring/decimal"
)

// Quote quotes a market order.
func (ob *OrderBook) Quote(traderID string, side Side, amount decimal.Decimal) (*Quote, error) {
	defer ob.RUnlock()
	ob.RLock()

	if strings.TrimSpace(traderID) == "" {
		return nil, ErrInvalidTraderID
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return nil, ErrInvalidAmount
	}

	price := decimal.Zero

	var (
		level *OrderQueue
		next  func(decimal.Decimal) *OrderQueue
	)

	if side == Buy {
		level = ob.asks.MinPriceQueue()
		next = ob.asks.GreaterThan
	} else {
		level = ob.bids.MaxPriceQueue()
		next = ob.bids.LessThan
	}

	for level != nil && amount.GreaterThan(decimal.Zero) {
		headOrderEl := level.Front()

		for headOrderEl != nil && amount.GreaterThan(decimal.Zero) {
			headOrder := headOrderEl.Value.(*Order)

			if headOrder.traderID == traderID {
				headOrderEl = headOrderEl.Next()
				continue
			}

			if amount.GreaterThanOrEqual(headOrder.amount) {
				price = price.Add(headOrder.price.Mul(headOrder.amount))
				amount = amount.Sub(headOrder.amount)
			} else {
				price = price.Add(headOrder.price.Mul(amount))
				amount = decimal.Zero
			}

			headOrderEl = headOrderEl.Next()
		}

		level = next(level.price)
	}

	return NewQuote(price, amount), nil
}
