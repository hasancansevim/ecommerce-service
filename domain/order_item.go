package domain

type OrderItem struct {
	Id        int64
	OrderId   int64
	ProductId int64
	Quantity  int
	Price     float32
}
