package orderbook

import (
	"strings"

	"github.com/shopspring/decimal"
)

// ProcessMarketOrder processes a market order.
func (ob *OrderBook) ProcessMarketOrder(orderID, traderID string, side Side, amount, price decimal.Decimal) ([]*Trade, error) {
	defer func() {
		ob.version++
		ob.Unlock()
	}()

	ob.Lock()

	if strings.TrimSpace(orderID) == "" {
		return nil, ErrInvalidOrderID
	}

	if ob.orders[orderID] != nil {
		return nil, ErrOrderAlreadyExists
	}

	if strings.TrimSpace(traderID) == "" {
		return nil, ErrInvalidTraderID
	}

	if amount.LessThanOrEqual(decimal.Zero) {
		return nil, ErrInvalidAmount
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return nil, ErrInvalidPrice
	}

	var (
		sideToProcess *OrderSide
		level         *OrderQueue
		next          func(decimal.Decimal) *OrderQueue
	)

	if side == Buy {
		sideToProcess = ob.asks
		level = ob.asks.MinPriceQueue()
		next = ob.asks.GreaterThan
	} else {
		sideToProcess = ob.bids
		level = ob.bids.MaxPriceQueue()
		next = ob.bids.LessThan
	}

	amountToTrade := amount
	priceToTrade := price
	trades := make([]*Trade, 0)

	for level != nil && amountToTrade.GreaterThan(decimal.Zero) && priceToTrade.GreaterThan(decimal.Zero) {
		headOrderEl := level.Front()

		for headOrderEl != nil && amountToTrade.GreaterThan(decimal.Zero) && priceToTrade.GreaterThan(decimal.Zero) {
			headOrder := headOrderEl.Value.(*Order)

			if headOrder.traderID == traderID {
				headOrderEl = headOrderEl.Next()
				continue
			}

			if amount.GreaterThanOrEqual(headOrder.amount) {
				trades = append(trades, NewTrade(orderID, headOrder.id, headOrder.amount, headOrder.price))
				amountToTrade = amountToTrade.Sub(headOrder.amount)
				priceToTrade = priceToTrade.Sub(headOrder.price)

				headOrderEl = headOrderEl.Next()
				ob.remove(headOrder.id)
			} else {
				trades = append(trades, NewTrade(orderID, headOrder.id, amountToTrade, headOrder.price))
				sideToProcess.UpdateAmount(headOrderEl, headOrder.amount.Sub(amountToTrade))
				amountToTrade = decimal.Zero
				priceToTrade = decimal.Zero

				headOrderEl = headOrderEl.Next()
			}
		}

		level = next(level.price)
	}

	return trades, nil
}
