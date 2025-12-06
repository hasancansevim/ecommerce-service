package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
)

type ICartItemService interface {
	AddItemToCart(cartItem dto.CreateCartItemRequest) (dto.CartItemResponse, error)
	GetItemsByCartId(cartId int64) []dto.CartItemResponse
	UpdateItemQuantity(cartItemId int64, newQuantity int) (dto.CartItemResponse, error)
	RemoveItemFromCart(cartItemId int64) error
	ClearCartItems(cartId int64) error
	IncreaseItemQuantity(cartItemId int64, amount int) error
	DecreaseItemQuantity(cartItemId int64, amount int) error
}

type CartItemService struct {
	cartItemRepository persistence.ICartItemRepository
}

func NewCartItemService(cartItemRepository persistence.ICartItemRepository) ICartItemService {
	return &CartItemService{
		cartItemRepository: cartItemRepository,
	}
}

func (cartItemService *CartItemService) AddItemToCart(cartItem dto.CreateCartItemRequest) (dto.CartItemResponse, error) {
	item, err := cartItemService.cartItemRepository.AddItemToCart(domain.CartItem{
		CartId:    cartItem.CartId,
		ProductId: cartItem.ProductId,
		Quantity:  cartItem.Quantity,
	})

	if err != nil {
		return dto.CartItemResponse{}, _errors.NewBadRequest(err.Error())
	}

	cartItemDto := dto.CartItemResponse{
		CartId:    item.CartId,
		ProductId: item.ProductId,
		Quantity:  item.Quantity,
	}
	return cartItemDto, nil
}

func (cartItemService *CartItemService) GetItemsByCartId(cartId int64) []dto.CartItemResponse {
	items := cartItemService.cartItemRepository.GetItemsByCartId(cartId)
	itemsDto := make([]dto.CartItemResponse, 0, len(items))
	for _, item := range items {
		itemsDto = append(itemsDto, dto.CartItemResponse{
			CartId:    item.CartId,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		})
	}
	return itemsDto
}

func (cartItemService *CartItemService) UpdateItemQuantity(cartItemId int64, newQuantity int) (dto.CartItemResponse, error) {
	item, err := cartItemService.cartItemRepository.UpdateItemQuantity(cartItemId, newQuantity)
	if err != nil {
		return dto.CartItemResponse{}, _errors.NewBadRequest(err.Error())
	}
	cartItemDto := dto.CartItemResponse{
		CartId:    item.CartId,
		ProductId: item.ProductId,
		Quantity:  item.Quantity,
	}
	return cartItemDto, nil
}

func (cartItemService *CartItemService) RemoveItemFromCart(cartItemId int64) error {
	return cartItemService.cartItemRepository.RemoveItemFromCart(cartItemId)
}

func (cartItemService *CartItemService) ClearCartItems(cartId int64) error {
	return cartItemService.cartItemRepository.ClearCartItems(cartId)
}

func (cartItemService *CartItemService) IncreaseItemQuantity(cartItemId int64, amount int) error {
	return cartItemService.cartItemRepository.IncreaseItemQuantity(cartItemId, amount)
}

func (cartItemService *CartItemService) DecreaseItemQuantity(cartItemId int64, amount int) error {
	return cartItemService.cartItemRepository.DecreaseItemQuantity(cartItemId, amount)
}
