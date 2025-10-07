package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type ICartItemRepository interface {
	AddItemToCart(cartItem domain.CartItem) error
	UpdateItemQuantity(cart_item_id int64, newQuantity int) error
	RemoveItemFromCart(cart_item_id int64) error
	GetItemsByCartId(cart_id int64) []domain.CartItem
	ClearCartItems(cart_id int64) error
	IncreaseItemQuantity(cart_item_id int64, amount int) error
	DecreaseItemQuantity(cart_item_id int64, amount int) error
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

func (cartItemRepository *CartItemRepository) AddItemToCart(cartItem domain.CartItem) error {
	ctx := context.Background()
	err := cartItemRepository.scanner.ExecuteExec(ctx, "insert into cart_items (cart_id,product_id,quantity) values ($1,$2,$3)",
		cartItem.CartId, cartItem.ProductId, cartItem.Quantity)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (cartItemRepository *CartItemRepository) UpdateItemQuantity(cart_item_id int64, newQuantity int) error {
	ctx := context.Background()
	err := cartItemRepository.scanner.ExecuteExec(ctx, "update cart_items set quantity=$1 where id=$2", newQuantity, cart_item_id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (cartItemRepository *CartItemRepository) RemoveItemFromCart(cart_item_id int64) error {
	ctx := context.Background()
	err := cartItemRepository.scanner.ExecuteExec(ctx, "elete from cart_items where id=$1", cart_item_id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (cartItemRepository *CartItemRepository) GetItemsByCartId(cart_id int64) []domain.CartItem {
	ctx := context.Background()
	items, err := cartItemRepository.scanner.QueryAndScan(ctx, "select * from cart_items where cart_id = $1", cart_id)
	if err != nil {
		log.Error(err)
		return []domain.CartItem{}
	}
	return items
}

func (cartItemRepository *CartItemRepository) ClearCartItems(cart_id int64) error {
	ctx := context.Background()
	err := cartItemRepository.scanner.ExecuteExec(ctx, "delete from cart_items where cart_id=$1", cart_id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (cartItemRepository *CartItemRepository) IncreaseItemQuantity(cart_item_id int64, amount int) error {
	ctx := context.Background()
	item, err := cartItemRepository.scanner.QueryRowAndScan(ctx, "select * from cart_items where id = $1", cart_item_id)
	if err != nil {
		log.Error(err)
		return err
	}
	increasedQuantity := item.Quantity + amount
	executeExecErr := cartItemRepository.scanner.ExecuteExec(ctx, "update cart_items set quantity=$1 where id=$2", increasedQuantity, cart_item_id)
	if executeExecErr != nil {
		log.Error(executeExecErr)
		return executeExecErr
	}
	return nil
}

func (cartItemRepository *CartItemRepository) DecreaseItemQuantity(cart_item_id int64, amount int) error {
	ctx := context.Background()
	item, err := cartItemRepository.scanner.QueryRowAndScan(ctx, "select * from cart_items where id = $1", cart_item_id)
	if err != nil {
		log.Error(err)
		return err
	}
	decreasedQuantity := item.Quantity - amount
	executeExecErr := cartItemRepository.scanner.ExecuteExec(ctx, "update cart_items set quantity=$1 where id=$2", decreasedQuantity, cart_item_id)
	if executeExecErr != nil {
		log.Error(executeExecErr)
		return executeExecErr
	}
	return nil
}
