package domain

type CartItem struct {
	Id        int64 `json:"id"`
	CartId    int64 `json:"cart_id"`
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
