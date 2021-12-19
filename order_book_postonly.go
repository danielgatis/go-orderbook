package orderbook

import (
	"strings"

	"github.com/shopspring/decimal"
)

// ProcessPostOnlyOrder processes a post only order.
func (ob *OrderBook) ProcessPostOnlyOrder(orderID, traderID string, side Side, amount, price decimal.Decimal) ([]*Trade, error) {
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

	order := NewOrder(orderID, traderID, side, amount, price)

	if side == Buy {
		ob.orders[order.id] = ob.bids.Append(order)
	} else {
		ob.orders[order.id] = ob.asks.Append(order)
	}

	return make([]*Trade, 0), nil
}
