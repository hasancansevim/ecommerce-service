package persistence

import (
	"context"
	"errors"
	"fmt"
	"go-ecommerce-service/domain"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IOrderItemRepository interface {
	AddOrderItem(orderItem domain.OrderItem) error
	GetOrderItemById(orderItem_id int64) (domain.OrderItem, error)
	GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error)
	GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error)
	UpdateOrderItem(orderItem_id int64, orderItem domain.OrderItem) error
	UpdateOrderItemQuantity(orderItem_id int64, quantity int) error
	DeleteOrderItemById(orderItem_id int64) error
	DeleteAllOrderItemsByOrderId(orderId int64) error
}

type OrderItemRepository struct {
	dbPool *pgxpool.Pool
}

func NewOrderItemRepository(dbPool *pgxpool.Pool) IOrderItemRepository {
	return &OrderItemRepository{
		dbPool: dbPool,
	}
}

func (orderItem *OrderItemRepository) AddOrderItem(order domain.OrderItem) error {
	ctx := context.Background()
	addOrderItemSql := `insert into order_items (order_id, product_id, quantity, price) values($1,$2,$3,$4)`
	_, execErr := orderItem.dbPool.Exec(ctx, addOrderItemSql, order.OrderId, order.ProductId, order.Quantity, order.Price)
	if execErr != nil {
		return execErr
	}
	return nil
}

func (orderItem *OrderItemRepository) GetOrderItemById(orderItem_id int64) (domain.OrderItem, error) {
	ctx := context.Background()
	getOrderItemById := `select * from order_items where id = $1`
	queryRow := orderItem.dbPool.QueryRow(ctx, getOrderItemById, orderItem_id)

	if queryRow == nil {
		return domain.OrderItem{}, errors.New(fmt.Sprintf("Error while something went wrong in GetOrderItemById"))
	}
	return extractOrderItemFromRow(queryRow), nil
}

func (orderItem *OrderItemRepository) GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error) {
	ctx := context.Background()
	getOrderItemByOrderId := `select * from order_items where order_id = $1`
	queryRows, queryRowsErr := orderItem.dbPool.Query(ctx, getOrderItemByOrderId, orderId)
	if queryRowsErr != nil {
		log.Error(queryRowsErr)
		return []domain.OrderItem{}, nil
	}
	orderItemsFromQueryRows, orderItemsFromQueryRowsErr := extractOrderItemsFromQueryRows(queryRows)
	if orderItemsFromQueryRowsErr != nil {
		log.Error(orderItemsFromQueryRowsErr)
		return []domain.OrderItem{}, nil
	}

	return orderItemsFromQueryRows, nil
}

func (orderItem *OrderItemRepository) GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error) {
	ctx := context.Background()
	getOrderItemById := `select * from order_items where product_id = $1`
	queryRows, queryRowsErr := orderItem.dbPool.Query(ctx, getOrderItemById, productId)
	if queryRowsErr != nil {
		log.Error(queryRowsErr)
		return []domain.OrderItem{}, nil
	}
	orderItemsFromQueryRows, orderItemsFromQueryRowsErr := extractOrderItemsFromQueryRows(queryRows)
	if orderItemsFromQueryRowsErr != nil {
		log.Error(orderItemsFromQueryRowsErr)
		return []domain.OrderItem{}, nil
	}
	return orderItemsFromQueryRows, nil
}

func (orderItem *OrderItemRepository) UpdateOrderItem(orderItem_id int64, order_item domain.OrderItem) error {
	ctx := context.Background()
	updateOrderItemSql := `update order_items set order_id=$1,product_id=$2,quantity=$3,price=$4 where id=$5`
	_, execErr := orderItem.dbPool.Exec(ctx, updateOrderItemSql, order_item.OrderId, order_item.ProductId, order_item.Quantity, order_item.Price, order_item.Id)
	if execErr != nil {
		return execErr
	}
	return nil
}

func (orderItem *OrderItemRepository) UpdateOrderItemQuantity(orderItem_id int64, quantity int) error {
	ctx := context.Background()
	updateOrderItemSql := `update order_items set quantity=$1 where id=$2`
	_, execErr := orderItem.dbPool.Exec(ctx, updateOrderItemSql, quantity, orderItem_id)
	if execErr != nil {
		return execErr
	}
	return nil
}

func (orderItem *OrderItemRepository) DeleteOrderItemById(orderItem_id int64) error {
	ctx := context.Background()
	deleteOrderItemByIdSql := `delete from order_items where id = $1`
	_, execErr := orderItem.dbPool.Exec(ctx, deleteOrderItemByIdSql, orderItem_id)
	if execErr != nil {
		return execErr
	}
	return nil
}

func (orderItem *OrderItemRepository) DeleteAllOrderItemsByOrderId(orderId int64) error {
	ctx := context.Background()
	deleteAllOrderItemsSql := `delete from order_items where order_id = $1`
	_, execErr := orderItem.dbPool.Exec(ctx, deleteAllOrderItemsSql, orderId)
	if execErr != nil {
		return execErr
	}
	return nil
}

func extractOrderItemFromRow(row pgx.Row) domain.OrderItem {
	var orderItem domain.OrderItem
	var id int64
	var order_id int64
	var product_id int64
	var quantity int
	var price float32

	row.Scan(&id, &order_id, &product_id, &quantity, &price)

	orderItem = domain.OrderItem{
		Id:        id,
		OrderId:   order_id,
		ProductId: product_id,
		Quantity:  quantity,
		Price:     price,
	}
	return orderItem
}

func extractOrderItemsFromQueryRows(rows pgx.Rows) ([]domain.OrderItem, error) {
	var orderItems []domain.OrderItem
	var id int64
	var order_id int64
	var product_id int64
	var quantity int
	var price float32

	for rows.Next() {
		rows.Scan(&id, &order_id, &product_id, &quantity, &price)
		orderItems = append(orderItems, domain.OrderItem{
			Id:        id,
			OrderId:   order_id,
			ProductId: product_id,
			Quantity:  quantity,
			Price:     price,
		})
	}
	return orderItems, nil
}
