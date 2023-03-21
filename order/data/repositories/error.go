package repositories

import "errors"

var OrderNotFound = errors.New("Order with the given id, not found!!")

var OrderStatusAlreadyUpdated = errors.New("Order status already updated!")
