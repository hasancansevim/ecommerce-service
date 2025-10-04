package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type ICartRepository interface {
	GetCartsByUserId(userId int64) []domain.Cart
	GetCartById(cartId int64) domain.Cart
	CreateCart(cart domain.Cart) error
	DeleteCartById(cartId int64) error
	ClearUserCart(userId int64) error
}

type CartRepository struct {
	dbPool *pgxpool.Pool
}

func NewCartRepository(dbPool *pgxpool.Pool) ICartRepository {
	return &CartRepository{
		dbPool: dbPool,
	}
}

func (cartRepository *CartRepository) GetCartsByUserId(userId int64) []domain.Cart {
	ctx := context.Background()
	getCartByUserIdSql := `select * from carts where user_id = $1`

	queryRows, queryRowsErr := cartRepository.dbPool.Query(ctx, getCartByUserIdSql, userId)

	if queryRowsErr != nil {
		log.Error(queryRowsErr)
		return []domain.Cart{}
	}

	return extractCartFromRows(queryRows)
}

func (cartRepository *CartRepository) GetCartById(cartId int64) domain.Cart {
	ctx := context.Background()
	getCartByIdSql := `select * from carts where id = $1`

	query, queryErr := cartRepository.dbPool.Query(ctx, getCartByIdSql, cartId)

	if queryErr != nil {
		log.Error(queryErr)
		return domain.Cart{}
	}

	var cart domain.Cart
	var id int64
	var user_id int64
	var created_at time.Time
	for query.Next() {
		query.Scan(&id, &user_id, &created_at)
		cart = domain.Cart{
			Id:        id,
			UserId:    user_id,
			CreatedAt: created_at,
		}
	}
	return cart
}

func (cartRepository *CartRepository) CreateCart(cart domain.Cart) error {
	ctx := context.Background()
	createCartSql := `insert into carts (user_id, created_at) values ($1, $2)`
	_, cartExecErr := cartRepository.dbPool.Exec(ctx, createCartSql, cart.UserId, time.Now())
	if cartExecErr != nil {
		log.Error(cartExecErr)
		return cartExecErr
	}
	return nil
}

func (cartRepository *CartRepository) DeleteCartById(cartId int64) error {
	ctx := context.Background()
	deleteCartByIdSql := `delete from carts where id = $1`
	_, execErr := cartRepository.dbPool.Exec(ctx, deleteCartByIdSql, cartId)
	if execErr != nil {
		log.Error(execErr)
		return execErr
	}
	return nil
}

func (cartRepository *CartRepository) ClearUserCart(userId int64) error {
	ctx := context.Background()
	clearUserCartSql := `delete from carts where user_id = $1`
	_, execErr := cartRepository.dbPool.Exec(ctx, clearUserCartSql, userId)
	if execErr != nil {
		log.Error(execErr)
		return execErr
	}
	return nil

}
func extractCartFromRows(rows pgx.Rows) []domain.Cart {
	var carts []domain.Cart

	var id int64
	var user_id int64
	var created_at time.Time

	for rows.Next() {
		rows.Scan(&id, &user_id, &created_at)
		carts = append(carts, domain.Cart{
			Id:        id,
			UserId:    user_id,
			CreatedAt: created_at,
		})
	}
	return carts
}
