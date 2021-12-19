package orderbook

import "errors"

// Orderbook erros
var (
	ErrInvalidOrderID     = errors.New("Invalid order id")
	ErrInvalidTraderID    = errors.New("Invalid trader id")
	ErrInvalidAmount      = errors.New("Invalid amount")
	ErrInvalidPrice       = errors.New("Invalid price")
	ErrInvalidSide        = errors.New("Invalid side")
	ErrOrderAlreadyExists = errors.New("Order already exists")
)
