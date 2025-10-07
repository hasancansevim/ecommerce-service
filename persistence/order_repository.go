package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type IOrderRepository interface {
	CreateOrder(order domain.Order) error
	GetOrderById(orderId int64) domain.Order
	GetOrdersByUserId(userId int64) ([]domain.Order, error)
	GetAllOrders() ([]domain.Order, error)
	UpdateOrderStatus(orderId int64, status bool) error
	DeleteOrderById(orderId int64) error
	UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) error
	GetOrdersByStatus(status string) ([]domain.Order, error)
}

type OrderRepository struct {
	dbPool  *pgxpool.Pool
	scanner *helper.GenericScanner[domain.Order]
}

func NewOrderRepository(dbPool *pgxpool.Pool) IOrderRepository {
	return &OrderRepository{
		dbPool:  dbPool,
		scanner: helper.NewGenericScanner(dbPool, helper.ScanOrder),
	}
}

func (orderRepository *OrderRepository) CreateOrder(order domain.Order) error {
	ctx := context.Background()
	err := orderRepository.scanner.ExecuteExec(ctx, "insert into orders (user_id,total_price,status) values ($1,$2,$3)",
		order.UserId, order.TotalPrice, order.Status)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) GetOrderById(orderId int64) domain.Order {
	ctx := context.Background()
	order, err := orderRepository.scanner.QueryRowAndScan(ctx, "select * from orders where id = $1", orderId)
	if err != nil {
		log.Error(err)
		return domain.Order{}
	}
	return order
}

func (orderRepository *OrderRepository) GetOrdersByUserId(userId int64) ([]domain.Order, error) {
	ctx := context.Background()
	orders, err := orderRepository.scanner.QueryAndScan(ctx, "select * from orders where user_id = $1", userId)
	if err != nil {
		log.Error(err)
		return []domain.Order{}, err
	}
	return orders, nil
}

func (orderRepository *OrderRepository) GetAllOrders() ([]domain.Order, error) {
	ctx := context.Background()
	orders, err := orderRepository.scanner.QueryAndScan(ctx, "select * from orders")
	if err != nil {
		log.Error(err)
		return []domain.Order{}, err
	}
	return orders, nil
}

func (orderRepository *OrderRepository) UpdateOrderStatus(orderId int64, status bool) error {
	ctx := context.Background()
	err := orderRepository.scanner.ExecuteExec(ctx, "update orders set status = $1 where id = $2", status, orderId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) DeleteOrderById(orderId int64) error {
	ctx := context.Background()
	err := orderRepository.scanner.ExecuteExec(ctx, "delete from orders where id = $1", orderId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) error {
	ctx := context.Background()
	err := orderRepository.scanner.ExecuteExec(ctx, "update orders set total_price = $1 where id = $2", newTotalPrice, orderId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) GetOrdersByStatus(status string) ([]domain.Order, error) {
	ctx := context.Background()
	orders, err := orderRepository.scanner.QueryAndScan(ctx, "select * from orders where status = $1", status)
	if err != nil {
		log.Error(err)
		return []domain.Order{}, err
	}
	return orders, nil
}
