package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type ICartRepository interface {
	GetCartsByUserId(userId int64) []domain.Cart
	GetCartById(cartId int64) domain.Cart
	CreateCart(cart domain.Cart) (domain.Cart, error)
	DeleteCartById(cartId int64) error
	ClearUserCart(userId int64) error
}

type CartRepository struct {
	dbPool  *pgxpool.Pool
	scanner *helper.GenericScanner[domain.Cart]
}

func NewCartRepository(dbPool *pgxpool.Pool) ICartRepository {
	return &CartRepository{
		dbPool:  dbPool,
		scanner: helper.NewGenericScanner(dbPool, helper.ScanCart),
	}
}

func (cartRepository *CartRepository) GetCartsByUserId(userId int64) []domain.Cart {
	ctx := context.Background()
	carts, err := cartRepository.scanner.QueryAndScan(ctx, "select * from carts where user_id = $1", userId)
	if err != nil {
		log.Error(err)
		return []domain.Cart{}
	}
	return carts
}

func (cartRepository *CartRepository) GetCartById(cartId int64) domain.Cart {
	ctx := context.Background()
	cart, err := cartRepository.scanner.QueryRowAndScan(ctx, "select * from carts where id = $1", cartId)
	if err != nil {
		log.Error(err)
		return domain.Cart{}
	}
	return cart
}

func (cartRepository *CartRepository) CreateCart(cart domain.Cart) (domain.Cart, error) {
	ctx := context.Background()
	query := `INSERT INTO carts (user_id, created_at) VALUES ($1, $2) RETURNING *`
	cart, err := cartRepository.scanner.QueryRowAndScan(ctx, query, cart.UserId, cart.CreatedAt)
	if err != nil {
		return domain.Cart{}, err
	}
	return cart, nil
}

func (cartRepository *CartRepository) DeleteCartById(cartId int64) error {
	ctx := context.Background()
	err := cartRepository.scanner.ExecuteExec(ctx, "delete from carts where id = $1", cartId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (cartRepository *CartRepository) ClearUserCart(userId int64) error {
	ctx := context.Background()
	err := cartRepository.scanner.ExecuteExec(ctx, "delete from carts where user_id = $1", userId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
