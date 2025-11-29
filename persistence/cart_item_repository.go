package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type ICartItemRepository interface {
	AddItemToCart(cartItem domain.CartItem) (domain.CartItem, error)
	UpdateItemQuantity(cartItemId int64, newQuantity int) (domain.CartItem, error)
	RemoveItemFromCart(cartItemId int64) error
	GetItemsByCartId(cartId int64) []domain.CartItem
	ClearCartItems(cartId int64) error
	IncreaseItemQuantity(cartItemId int64, amount int) error
	DecreaseItemQuantity(cartItemId int64, amount int) error
}

type CartItemRepository struct {
	dbPool  *pgxpool.Pool
	scanner *helper.GenericScanner[domain.CartItem]
}

func NewCartItemRepository(dbPool *pgxpool.Pool) ICartItemRepository {
	return &CartItemRepository{
		dbPool:  dbPool,
		scanner: helper.NewGenericScanner(dbPool, helper.ScanCartItem),
	}
}

func (cartItemRepository *CartItemRepository) AddItemToCart(cartItem domain.CartItem) (domain.CartItem, error) {
	ctx := context.Background()
	query := `INSERT INTO cart_items (cart_id,product_id,quantity) VALUES ($1,$2,$3) RETURNING *`
	cartItem, err := cartItemRepository.scanner.QueryRowAndScan(ctx, query,
		cartItem.CartId, cartItem.ProductId, cartItem.Quantity)
	if err != nil {
		return domain.CartItem{}, err
	}
	return cartItem, nil
}

func (cartItemRepository *CartItemRepository) UpdateItemQuantity(cartItemId int64, newQuantity int) (domain.CartItem, error) {
	ctx := context.Background()
	query := `UPDATE cart_items set quantity=$1 where id=$2 RETURNING *`
	cartItem, err := cartItemRepository.scanner.QueryRowAndScan(ctx, query, newQuantity, cartItemId)
	if err != nil {
		return domain.CartItem{}, err
	}
	return cartItem, nil
}

func (cartItemRepository *CartItemRepository) RemoveItemFromCart(cartItemId int64) error {
	ctx := context.Background()
	query := `SELECT from cart_items where id=$1`
	err := cartItemRepository.scanner.ExecuteExec(ctx, query, cartItemId)
	if err != nil {
		return err
	}
	return nil
}

func (cartItemRepository *CartItemRepository) GetItemsByCartId(cartId int64) []domain.CartItem {
	ctx := context.Background()
	query := `SELECT * from cart_items where cart_id = $1`
	items, err := cartItemRepository.scanner.QueryAndScan(ctx, query, cartId)
	if err != nil {
		log.Error(err)
		return []domain.CartItem{}
	}
	return items
}

func (cartItemRepository *CartItemRepository) ClearCartItems(cartId int64) error {
	ctx := context.Background()
	query := `DELETE from cart_items where cart_id=$1`
	err := cartItemRepository.scanner.ExecuteExec(ctx, query, cartId)
	if err != nil {
		return err
	}
	return nil
}

func (cartItemRepository *CartItemRepository) IncreaseItemQuantity(cartItemId int64, amount int) error {
	ctx := context.Background()
	query := `SELECT * from cart_items where id = $1`
	item, err := cartItemRepository.scanner.QueryRowAndScan(ctx, query, cartItemId)
	if err != nil {
		return err
	}
	increasedQuantity := item.Quantity + amount

	execQuery := `UPDATE cart_items set quantity=$1 where id=$2`
	executeExecErr := cartItemRepository.scanner.ExecuteExec(ctx, execQuery, increasedQuantity, cartItemId)
	if executeExecErr != nil {
		return executeExecErr
	}
	return nil
}

func (cartItemRepository *CartItemRepository) DecreaseItemQuantity(cartItemId int64, amount int) error {
	ctx := context.Background()
	query := `SELECT * from cart_items where id = $1`
	item, err := cartItemRepository.scanner.QueryRowAndScan(ctx, query, cartItemId)
	if err != nil {
		return err
	}
	decreasedQuantity := item.Quantity - amount
	execQuery := `UPDATE cart_items set quantity=$1 where id=$2`
	executeExecErr := cartItemRepository.scanner.ExecuteExec(ctx, execQuery, decreasedQuantity, cartItemId)
	if executeExecErr != nil {
		return executeExecErr
	}
	return nil
}
