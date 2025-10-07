package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

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
	dbPool  *pgxpool.Pool
	scanner *helper.GenericScanner[domain.OrderItem]
}

func NewOrderItemRepository(dbPool *pgxpool.Pool) IOrderItemRepository {
	return &OrderItemRepository{
		dbPool:  dbPool,
		scanner: helper.NewGenericScanner(dbPool, helper.ScanOrderItem),
	}
}

func (orderItem *OrderItemRepository) AddOrderItem(order_item domain.OrderItem) error {
	ctx := context.Background()
	err := orderItem.scanner.ExecuteExec(ctx, "insert into order_items (order_id, product_id, quantity, price) values($1,$2,$3,$4)",
		order_item.Id, order_item.ProductId, order_item.Quantity, order_item.Price)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderItem *OrderItemRepository) GetOrderItemById(orderItem_id int64) (domain.OrderItem, error) {
	ctx := context.Background()
	item, err := orderItem.scanner.QueryRowAndScan(ctx, "select * from order_items where id = $1", orderItem_id)
	if err != nil {
		log.Error(err)
		return domain.OrderItem{}, err
	}
	return item, nil
}

func (orderItem *OrderItemRepository) GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error) {
	ctx := context.Background()
	orderItems, err := orderItem.scanner.QueryAndScan(ctx, "select * from order_items where order_id = $1", orderId)
	if err != nil {
		log.Error(err)
		return []domain.OrderItem{}, err
	}
	return orderItems, nil
}

func (orderItem *OrderItemRepository) GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error) {
	ctx := context.Background()
	orderItems, err := orderItem.scanner.QueryAndScan(ctx, "select * from order_items where product_id = $1", productId)
	if err != nil {
		log.Error(err)
		return []domain.OrderItem{}, err
	}
	return orderItems, nil
}

func (orderItem *OrderItemRepository) UpdateOrderItem(orderItem_id int64, order_item domain.OrderItem) error {
	ctx := context.Background()
	err := orderItem.scanner.ExecuteExec(ctx, "update order_items set order_id=$1,product_id=$2,quantity=$3,price=$4 where id=$5",
		order_item.OrderId, order_item.ProductId, order_item.Quantity, order_item.Price, order_item.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderItem *OrderItemRepository) UpdateOrderItemQuantity(orderItem_id int64, quantity int) error {
	ctx := context.Background()
	err := orderItem.scanner.ExecuteExec(ctx, "update order_items set quantity=$1 where id=$2", quantity, orderItem_id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderItem *OrderItemRepository) DeleteOrderItemById(orderItem_id int64) error {
	ctx := context.Background()
	err := orderItem.scanner.ExecuteExec(ctx, "delete from order_items where id = $1", orderItem_id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderItem *OrderItemRepository) DeleteAllOrderItemsByOrderId(orderId int64) error {
	ctx := context.Background()
	err := orderItem.scanner.ExecuteExec(ctx, "delete from order_items where order_id = $1", orderId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
