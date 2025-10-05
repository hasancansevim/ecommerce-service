package common

import (
	"errors"
	"fmt"
)

var (
	ErrProductNotFound = errors.New("Product not found")
	ErrDatabaseQuery   = errors.New("Database query error")
	ErrDatabaseExecute = errors.New("Database execution error")
)

func WrapError(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", operation, err)
}
