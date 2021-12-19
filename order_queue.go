package orderbook

import (
	"container/list"

	"github.com/shopspring/decimal"
)

// OrderQueue represents a queue of orders.
type OrderQueue struct {
	amount decimal.Decimal
	price  decimal.Decimal
	orders *list.List
}

// NewOrderQueue creates a new order queue.
func NewOrderQueue(price decimal.Decimal) *OrderQueue {
	return &OrderQueue{decimal.Zero, price, list.New()}
}

// Price returns the queue price.
func (oq *OrderQueue) Price() decimal.Decimal {
	return oq.price
}

// Amount returns the queue amout.
func (oq *OrderQueue) Amount() decimal.Decimal {
	return oq.amount
}

// Orders returns the orders as a list.
func (oq *OrderQueue) Orders() *list.List {
	return oq.orders
}

// Len returns the length.
func (oq *OrderQueue) Len() int {
	return oq.orders.Len()
}

// Front returns the first order of the queue.
func (oq *OrderQueue) Front() *list.Element {
	return oq.orders.Front()
}

// Back returns the last order of the queue.
func (oq *OrderQueue) Back() *list.Element {
	return oq.orders.Back()
}

// Append appends an order.
func (oq *OrderQueue) Append(order *Order) *list.Element {
	oq.amount = oq.amount.Add(order.amount)
	return oq.orders.PushBack(order)
}

// UpdateAmount updates an order amount.
func (oq *OrderQueue) UpdateAmount(e *list.Element, amount decimal.Decimal) *Order {
	order := e.Value.(*Order)
	oq.amount = oq.amount.Sub(order.amount)
	oq.amount = oq.amount.Add(amount)
	order.amount = amount
	return order
}

// Remove removes an order.
func (oq *OrderQueue) Remove(e *list.Element) *Order {
	oq.amount = oq.amount.Sub(e.Value.(*Order).amount)
	return oq.orders.Remove(e).(*Order)
}
