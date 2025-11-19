package interfaces

import (
	"go-ecommerce-service/domain"

	"github.com/jackc/pgx/v4"
)

type Scannable interface {
	domain.Product | domain.User | domain.Cart | domain.CartItem | domain.Order | domain.OrderItem | domain.Category | domain.Store
}
type Scanner[T Scannable] interface {
	Scan(row pgx.Row) (T, error)
	ScanAll(rows pgx.Rows) ([]T, error)
}
