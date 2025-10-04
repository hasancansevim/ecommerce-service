package persistence

import (
	"context"
	"go-ecommerce-service/domain"

	"github.com/jackc/pgx/v4"
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
	dbPool *pgxpool.Pool
}

func NewCartItemRepository(dbPool *pgxpool.Pool) ICartItemRepository {
	return &CartItemRepository{
		dbPool: dbPool,
	}
}

func (cartItemRepository *CartItemRepository) AddItemToCart(cartItem domain.CartItem) error {
	ctx := context.Background()
	addItemToCartSql := `insert into cart_items (cart_id,product_id,quantity) values ($1,$2,$3)`
	cartExec, cartExecErr := cartItemRepository.dbPool.Exec(ctx, addItemToCartSql, cartItem.CartId, cartItem.ProductId, cartItem.Quantity)
	if cartExecErr != nil {
		log.Error("Error while add cart item : %v ", cartExecErr.Error())
		return cartExecErr
	}
	log.Info("Added cart item : %v ", cartExec)
	return nil
}

func (cartItemRepository *CartItemRepository) UpdateItemQuantity(cart_item_id int64, newQuantity int) error {
	ctx := context.Background()
	updateItemQuantitySql := `update cart_items set quantity=$1 where id=$2`
	updatedCartItem, execErr := cartItemRepository.dbPool.Exec(ctx, updateItemQuantitySql, newQuantity, cart_item_id)
	if execErr != nil {
		log.Error("Error while update cart item : %v ", execErr.Error())
		return execErr
	}
	log.Info("Updated cart item : %v ", updatedCartItem)
	return nil
}

func (cartItemRepository *CartItemRepository) RemoveItemFromCart(cart_item_id int64) error {
	ctx := context.Background()
	removeItemFromCartSql := `delete from cart_items where id=$1`
	_, execErr := cartItemRepository.dbPool.Exec(ctx, removeItemFromCartSql, cart_item_id)
	if execErr != nil {
		log.Error("Error while remove cart item : %v ", execErr.Error())
		return execErr
	}
	log.Info("Removed cart item : %v ", cart_item_id)
	return nil
}

func (cartItemRepository *CartItemRepository) GetItemsByCartId(cart_id int64) []domain.CartItem {
	ctx := context.Background()
	getItemsByCartIdSql := `select * from cart_items where cart_id = $1`
	queryRow, queryRowErr := cartItemRepository.dbPool.Query(ctx, getItemsByCartIdSql, cart_id)

	if queryRowErr != nil {
		log.Error("Error while get cart item : %v ", queryRowErr.Error())
		return []domain.CartItem{}
	}
	return extractCartItemsFromRows(queryRow)
}

func (cartItemRepository *CartItemRepository) ClearCartItems(cart_id int64) error {
	ctx := context.Background()
	clearCartItemsSql := `delete from cart_items where cart_id=$1`
	_, execErr := cartItemRepository.dbPool.Exec(ctx, clearCartItemsSql, cart_id)
	if execErr != nil {
		log.Error("Error while clear cart item : %v ", execErr.Error())
		return execErr
	}
	log.Info("Cleared cart items : %v ", cart_id)
	return nil
}

func (cartItemRepository *CartItemRepository) IncreaseItemQuantity(cart_item_id int64, amount int) error {
	ctx := context.Background()
	getCartItemSql := `select * from cart_items where id = $1`
	queryRow, queryRowErr := cartItemRepository.dbPool.Query(ctx, getCartItemSql, cart_item_id)
	if queryRowErr != nil {
		log.Error("Error while get cart item : %v ", queryRowErr.Error())
		return queryRowErr
	}
	rows := extractCartItemsFromRows(queryRow)
	quantity := rows[0].Quantity
	increasedQuantity := quantity + amount
	increaseItemQuantitySql := `update cart_items set quantity=$1 where id=$2`
	_, execErr := cartItemRepository.dbPool.Exec(ctx, increaseItemQuantitySql, increasedQuantity, cart_item_id)
	if execErr != nil {
		log.Error("Error while increase cart item : %v ", execErr.Error())
		return execErr
	}
	log.Info("Increased cart item : %v ", cart_item_id)
	return nil
}

func (cartItemRepository *CartItemRepository) DecreaseItemQuantity(cart_item_id int64, amount int) error {
	ctx := context.Background()
	getCartItemSql := `select * from cart_items where id = $1`
	queryRow, queryRowErr := cartItemRepository.dbPool.Query(ctx, getCartItemSql, cart_item_id)
	if queryRowErr != nil {
		log.Error("Error while get cart item : %v ", queryRowErr.Error())
		return queryRowErr
	}
	rows := extractCartItemsFromRows(queryRow)
	quantity := rows[0].Quantity
	decreasedQuantity := quantity - amount
	decreaseItemQuantitySql := `update cart_items set quantity=$1 where id=$2`
	_, execErr := cartItemRepository.dbPool.Exec(ctx, decreaseItemQuantitySql, decreasedQuantity, cart_item_id)
	if execErr != nil {
		log.Error("Error while decrease cart item : %v ", execErr.Error())
		return execErr
	}
	log.Info("Decreased cart item : %v ", cart_item_id)
	return nil
}

func extractCartItemsFromRows(row pgx.Rows) []domain.CartItem {
	var cartItems []domain.CartItem
	var id int64
	var cart_id int64
	var product_id int64
	var quantity int

	for row.Next() {
		row.Scan(&id, &cart_id, &product_id, &quantity)
		cartItems = append(cartItems, domain.CartItem{
			Id:        id,
			CartId:    cart_id,
			ProductId: product_id,
			Quantity:  quantity,
		})
	}
	return cartItems
}
