package common

import (
	"errors"
	"fmt"
)

var (
	ErrProductNotFound   = errors.New("Product not found")
	ErrUserNotFound      = errors.New("User not found")
	ErrOrderNotFound     = errors.New("Order not found")
	ErrOrderItemNotFound = errors.New("Order item not found")
	ErrCartNotFound      = errors.New("Cart not found")
	ErrCartItemNotFound  = errors.New("Cart item not found")
	ErrDatabaseQuery     = errors.New("Database query error")
	ErrDatabaseExecute   = errors.New("Database execution error")
)

func WrapError(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", operation, err)
}
