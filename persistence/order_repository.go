package persistence

import (
	"context"
	"errors"
	"fmt"
	"go-ecommerce-service/domain"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
	GetOrderCountByUserId(userId int64) (int, error)
}

type OrderRepository struct {
	dbPool *pgxpool.Pool
}

func (orderRepository *OrderRepository) CreateOrder(order domain.Order) error {
	ctx := context.Background()
	createOrderSql := `insert into orders (user_id,total_price,status) values ($1,$2,$3)`
	_, createOrderErr := orderRepository.dbPool.Exec(ctx, createOrderSql, order.UserId, order.TotalPrice, order.Status)
	if createOrderErr != nil {
		return createOrderErr
	}
	return nil
}

func NewOrderRepository(dbPool *pgxpool.Pool) IOrderRepository {
	return &OrderRepository{
		dbPool: dbPool,
	}
}

func (orderRepository *OrderRepository) GetOrderById(orderId int64) domain.Order {
	ctx := context.Background()
	getOrderByIdSql := `select * from orders where id = $1`
	queryRow := orderRepository.dbPool.QueryRow(ctx, getOrderByIdSql, orderId)

	return extractOrderFromRow(queryRow)
}

func (orderRepository *OrderRepository) GetOrdersByUserId(userId int64) ([]domain.Order, error) {
	ctx := context.Background()
	getOrdersByUserIdSql := `select * from orders where user_id = $1`
	queryRows, queryErr := orderRepository.dbPool.Query(ctx, getOrdersByUserIdSql, userId)
	if queryErr != nil {
		return []domain.Order{}, errors.New(fmt.Sprintf("Error while getting all orders by user id :%v", queryErr.Error()))
	}
	ordersFromRows := extractOrderFromRows(queryRows)
	return ordersFromRows, nil
}

func (orderRepository *OrderRepository) GetAllOrders() ([]domain.Order, error) {
	ctx := context.Background()
	getAllOrdersSql := `select * from orders`
	queryRows, queryErr := orderRepository.dbPool.Query(ctx, getAllOrdersSql)
	if queryErr != nil {
		return []domain.Order{}, queryErr
	}
	ordersFromRows := extractOrderFromRows(queryRows)
	return ordersFromRows, nil
}

func (orderRepository *OrderRepository) UpdateOrderStatus(orderId int64, status bool) error {
	ctx := context.Background()
	updateOrderStatusSql := `update orders set status = $1 where id = $2`
	_, err := orderRepository.dbPool.Exec(ctx, updateOrderStatusSql, status, orderId)
	if err != nil {
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) DeleteOrderById(orderId int64) error {
	ctx := context.Background()
	deleteOrderByIdSql := `delete from orders where id = $1`
	_, err := orderRepository.dbPool.Exec(ctx, deleteOrderByIdSql, orderId)
	if err != nil {
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) error {
	ctx := context.Background()
	updateOrderTotalPriceSql := `update orders set total_price = $1 where id = $2`
	_, err := orderRepository.dbPool.Exec(ctx, updateOrderTotalPriceSql, newTotalPrice, orderId)
	if err != nil {
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) GetOrdersByStatus(status string) ([]domain.Order, error) {
	ctx := context.Background()
	getOrdersByStatusSql := `select * from orders where status = $1`
	queryRows, queryErr := orderRepository.dbPool.Query(ctx, getOrdersByStatusSql, status)
	if queryErr != nil {
		return []domain.Order{}, queryErr
	}
	ordersFromRows := extractOrderFromRows(queryRows)
	return ordersFromRows, nil
}

func (orderRepository *OrderRepository) GetOrderCountByUserId(userId int64) (int, error) {
	ctx := context.Background()
	getOrderCountSql := `select count(*) from orders where user_id = $1`
	query, err := orderRepository.dbPool.Query(ctx, getOrderCountSql, userId)
	if err != nil {
		return 0, err
	}
	var count int
	scanErr := query.Scan(&count)
	if scanErr != nil {
		return 0, err
	}
	return count, nil
}

func extractOrderFromRow(row pgx.Row) domain.Order {
	var order domain.Order
	var id int64
	var user_id int64
	var total_price float32
	var status bool
	var created_at time.Time

	row.Scan(&id, &user_id, &total_price, &status, &created_at)

	order = domain.Order{
		Id:         id,
		UserId:     user_id,
		TotalPrice: total_price,
		Status:     status,
		CreatedAt:  created_at,
	}
	return order
}

func extractOrderFromRows(rows pgx.Rows) []domain.Order {
	var orders []domain.Order
	var id int64
	var user_id int64
	var total_price float32
	var status bool
	var created_at time.Time

	for rows.Next() {
		rows.Scan(&id, &user_id, &total_price, &status, &created_at)
		orders = append(orders, domain.Order{
			Id:         id,
			UserId:     user_id,
			TotalPrice: total_price,
			Status:     status,
			CreatedAt:  created_at,
		})
	}
	return orders
}
