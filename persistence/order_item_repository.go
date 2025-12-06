package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IOrderItemRepository interface {
	AddOrderItem(orderItem domain.OrderItem) (domain.OrderItem, error)
	GetOrderItemById(orderItemId int64) (domain.OrderItem, error)
	GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error)
	GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error)
	UpdateOrderItem(orderItemId int64, orderItem domain.OrderItem) (domain.OrderItem, error)
	UpdateOrderItemQuantity(orderItemId int64, quantity int) (domain.OrderItem, error)
	DeleteOrderItemById(orderItemId int64) error
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

func (orderItemRepository *OrderItemRepository) AddOrderItem(orderItem domain.OrderItem) (domain.OrderItem, error) {
	ctx := context.Background()
	query := `insert into order_items (order_id, product_id, quantity, price) values($1,$2,$3,$4) RETURNING *`
	addedOrderItem, err := orderItemRepository.scanner.QueryRowAndScan(ctx, query,
		orderItem.Id, orderItem.ProductId, orderItem.Quantity, orderItem.Price)
	if err != nil {
		return domain.OrderItem{}, err
	}
	return addedOrderItem, nil
}

func (orderItemRepository *OrderItemRepository) GetOrderItemById(orderItemId int64) (domain.OrderItem, error) {
	ctx := context.Background()
	orderItem, err := orderItemRepository.scanner.QueryRowAndScan(ctx, "select * from order_items where id = $1", orderItemId)
	if err != nil {
		return domain.OrderItem{}, err
	}
	return orderItem, nil
}

func (orderItemRepository *OrderItemRepository) GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error) {
	ctx := context.Background()
	orderItems, err := orderItemRepository.scanner.QueryAndScan(ctx, "select * from order_items where order_id = $1", orderId)
	if err != nil {
		return []domain.OrderItem{}, err
	}
	return orderItems, nil
}

func (orderItemRepository *OrderItemRepository) GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error) {
	ctx := context.Background()
	orderItems, err := orderItemRepository.scanner.QueryAndScan(ctx, "select * from order_items where product_id = $1", productId)
	if err != nil {
		return []domain.OrderItem{}, err
	}
	return orderItems, nil
}

func (orderItemRepository *OrderItemRepository) UpdateOrderItem(orderItemId int64, orderItem domain.OrderItem) (domain.OrderItem, error) {
	ctx := context.Background()
	query := `update order_items set order_id=$1,product_id=$2,quantity=$3,price=$4 where id=$5 RETURNING *`
	updatedOrderItem, err := orderItemRepository.scanner.QueryRowAndScan(ctx, query,
		orderItem.OrderId, orderItem.ProductId, orderItem.Quantity, orderItem.Price, orderItem.Id)
	if err != nil {
		return domain.OrderItem{}, err
	}
	return updatedOrderItem, nil
}

func (orderItemRepository *OrderItemRepository) UpdateOrderItemQuantity(orderItemId int64, quantity int) (domain.OrderItem, error) {
	ctx := context.Background()
	query := `update order_items set quantity=$1 where id=$2 RETURNING *`
	updatedOrderItem, err := orderItemRepository.scanner.QueryRowAndScan(ctx, query, quantity, orderItemId)
	if err != nil {
		return domain.OrderItem{}, err
	}
	return updatedOrderItem, nil
}

func (orderItemRepository *OrderItemRepository) DeleteOrderItemById(orderItem_id int64) error {
	ctx := context.Background()
	err := orderItemRepository.scanner.ExecuteExec(ctx, "delete from order_items where id = $1", orderItem_id)
	if err != nil {
		return err
	}
	return nil
}

func (orderItemRepository *OrderItemRepository) DeleteAllOrderItemsByOrderId(orderId int64) error {
	ctx := context.Background()
	err := orderItemRepository.scanner.ExecuteExec(ctx, "delete from order_items where order_id = $1", orderId)
	if err != nil {
		return err
	}
	return nil
}
