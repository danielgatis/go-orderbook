package orderbook

import (
	"container/list"
	"sort"

	"github.com/emirpasic/gods/examples/redblacktreeextended"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/shopspring/decimal"
)

// OrderSide represents all the prices about bids or asks.
type OrderSide struct {
	side   Side
	tree   *redblacktreeextended.RedBlackTreeExtended
	queue  map[decimal.Decimal]*OrderQueue
	amount decimal.Decimal
	size   int
	depth  int
}

// NewOrderSide creates a new order side.
func NewOrderSide(side Side) *OrderSide {
	tree := &redblacktreeextended.RedBlackTreeExtended{
		Tree: redblacktree.NewWith(func(a, b interface{}) int {
			return a.(decimal.Decimal).Cmp(b.(decimal.Decimal))
		}),
	}

	return &OrderSide{side, tree, make(map[decimal.Decimal]*OrderQueue), decimal.Zero, 0, 0}
}

// Append appends an order.
func (os *OrderSide) Append(order *Order) *list.Element {
	price := order.price

	priceQueue, ok := os.queue[price]
	if !ok {
		priceQueue = NewOrderQueue(price)
		os.queue[order.price] = priceQueue
		os.tree.Put(price, priceQueue)
		os.depth++
	}

	os.size++
	os.amount = os.amount.Add(order.amount)
	return priceQueue.Append(order)
}

// Remove removes an order.
func (os *OrderSide) Remove(e *list.Element) *Order {
	order := e.Value.(*Order)
	price := order.price

	priceQueue := os.queue[price]
	o := priceQueue.Remove(e)

	if priceQueue.Len() == 0 {
		delete(os.queue, price)
		os.tree.Remove(price)
		os.depth--
	}

	os.size--
	os.amount = os.amount.Sub(o.Amount())
	return o
}

// UpdateAmount updates an order amount.
func (os *OrderSide) UpdateAmount(e *list.Element, amount decimal.Decimal) *Order {
	order := e.Value.(*Order)
	price := order.price

	os.amount = os.amount.Sub(order.amount)
	os.amount = os.amount.Add(amount)

	priceQueue := os.queue[price]
	o := priceQueue.UpdateAmount(e, amount)

	return o
}

// MaxPriceQueue returns the order queue for the max price.
func (os *OrderSide) MaxPriceQueue() *OrderQueue {
	if os.depth <= 0 {
		return nil
	}

	if value, found := os.tree.GetMax(); found {
		return value.(*OrderQueue)
	}

	return nil
}

// MinPriceQueue returns the order queue for the min price.
func (os *OrderSide) MinPriceQueue() *OrderQueue {
	if os.depth <= 0 {
		return nil
	}

	if value, found := os.tree.GetMin(); found {
		return value.(*OrderQueue)
	}

	return nil
}

// LessThan returns the order queue for the price less than the given price.
func (os *OrderSide) LessThan(price decimal.Decimal) *OrderQueue {
	tree := os.tree.Tree
	node := tree.Root

	var floor *redblacktree.Node
	for node != nil {
		if tree.Comparator(price, node.Key) > 0 {
			floor = node
			node = node.Right
		} else {
			node = node.Left
		}
	}

	if floor != nil {
		return floor.Value.(*OrderQueue)
	}

	return nil
}

// GreaterThan returns the order queue for the price greater than the given price.
func (os *OrderSide) GreaterThan(price decimal.Decimal) *OrderQueue {
	tree := os.tree.Tree
	node := tree.Root

	var ceiling *redblacktree.Node
	for node != nil {
		if tree.Comparator(price, node.Key) < 0 {
			ceiling = node
			node = node.Left
		} else {
			node = node.Right
		}
	}

	if ceiling != nil {
		return ceiling.Value.(*OrderQueue)
	}

	return nil
}

// Orders return all the orders sorted by price. Desc when side is buy. Asc when side is sell.
func (os *OrderSide) Orders() []*Order {
	orders := make([]*Order, 0)

	for _, price := range os.queue {
		iter := price.Front()

		for iter != nil {
			orders = append(orders, iter.Value.(*Order))
			iter = iter.Next()
		}
	}

	if os.side == Buy {
		sort.Slice(orders[:], func(i, j int) bool {
			return orders[i].price.GreaterThan(orders[j].price)
		})
	} else {
		sort.Slice(orders[:], func(i, j int) bool {
			return orders[i].price.LessThan(orders[j].price)
		})
	}

	return orders
}
