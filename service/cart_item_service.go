package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
)

type ICartItemService interface {
	AddItemToCart(cartItem model.CartItemCreate) error
	GetItemsByCartId(cart_id int64) []domain.CartItem
	UpdateItemQuantity(cart_item_id int64, newQuantity int) error
	RemoveItemFromCart(cart_item_id int64) error
	ClearCartItems(cart_id int64) error
	IncreaseItemQuantity(cart_item_id int64, amount int) error
	DecreaseItemQuantity(cart_item_id int64, amount int) error
}

type CartItemService struct {
	cartItemRepository persistence.ICartItemRepository
}

func NewCartItemService(cartItemRepository persistence.ICartItemRepository) ICartItemService {
	return &CartItemService{
		cartItemRepository: cartItemRepository,
	}
}

func (cartItemService *CartItemService) AddItemToCart(cartItem model.CartItemCreate) error {
	err := cartItemService.cartItemRepository.AddItemToCart(domain.CartItem{
		CartId:    cartItem.CartId,
		ProductId: cartItem.ProductId,
		Quantity:  cartItem.Quantity,
	})
	if err != nil {
		return err
	}
	return nil
}

func (cartItemService *CartItemService) GetItemsByCartId(cart_id int64) []domain.CartItem {
	return cartItemService.cartItemRepository.GetItemsByCartId(cart_id)
}

func (cartItemService *CartItemService) UpdateItemQuantity(cart_item_id int64, newQuantity int) error {
	return cartItemService.cartItemRepository.UpdateItemQuantity(cart_item_id, newQuantity)
}

func (cartItemService *CartItemService) RemoveItemFromCart(cart_item_id int64) error {
	return cartItemService.cartItemRepository.RemoveItemFromCart(cart_item_id)
}

func (cartItemService *CartItemService) ClearCartItems(cart_id int64) error {
	return cartItemService.cartItemRepository.ClearCartItems(cart_id)
}

func (cartItemService *CartItemService) IncreaseItemQuantity(cart_item_id int64, amount int) error {
	return cartItemService.cartItemRepository.IncreaseItemQuantity(cart_item_id, amount)
}

func (cartItemService *CartItemService) DecreaseItemQuantity(cart_item_id int64, amount int) error {
	return cartItemService.cartItemRepository.DecreaseItemQuantity(cart_item_id, amount)
}
