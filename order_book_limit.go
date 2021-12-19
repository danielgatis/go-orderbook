package orderbook

import (
	"strings"

	"github.com/shopspring/decimal"
)

// ProcessLimitOrder processes a limit order.
func (ob *OrderBook) ProcessLimitOrder(orderID, traderID string, side Side, amount, price decimal.Decimal) ([]*Trade, error) {
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
		sideToAdd     *OrderSide
		sideToProcess *OrderSide
		comparator    func(decimal.Decimal) bool
		best          func() *OrderQueue
		next          func(decimal.Decimal) *OrderQueue
	)

	if side == Buy {
		sideToAdd = ob.bids
		sideToProcess = ob.asks
		comparator = price.GreaterThanOrEqual
		best = ob.asks.MinPriceQueue
		next = ob.asks.GreaterThan
	} else {
		sideToAdd = ob.asks
		sideToProcess = ob.bids
		comparator = price.LessThanOrEqual
		best = ob.bids.MaxPriceQueue
		next = ob.bids.LessThan
	}

	trades := make([]*Trade, 0)
	amountToTrade := amount
	bestPrice := best()

	for bestPrice != nil && amountToTrade.GreaterThan(decimal.Zero) && comparator(bestPrice.price) {
		headOrderEl := bestPrice.Front()
		bestPrice = next(bestPrice.price)

		for headOrderEl != nil && amountToTrade.GreaterThan(decimal.Zero) {
			headOrder := headOrderEl.Value.(*Order)

			if headOrder.traderID == traderID {
				headOrderEl = headOrderEl.Next()
				continue
			}

			if amountToTrade.GreaterThanOrEqual(headOrder.amount) {
				trades = append(trades, NewTrade(orderID, headOrder.id, headOrder.amount, headOrder.price))
				amountToTrade = amountToTrade.Sub(headOrder.amount)

				headOrderEl = headOrderEl.Next()
				ob.remove(headOrder.id)
			} else {
				trades = append(trades, NewTrade(orderID, headOrder.id, amountToTrade, headOrder.price))
				sideToProcess.UpdateAmount(headOrderEl, headOrder.amount.Sub(amountToTrade))
				amountToTrade = decimal.Zero

				headOrderEl = headOrderEl.Next()
			}
		}
	}

	if amountToTrade.GreaterThan(decimal.Zero) {
		order := NewOrder(orderID, traderID, side, amountToTrade, price)
		ob.orders[order.id] = sideToAdd.Append(order)
	}

	return trades, nil
}
