package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/service/validation"
)

type ICartService interface {
	GetCartsByUserId(userId int64) []domain.Cart
	GetCartById(cartId int64) domain.Cart
	CreateCart(cart model.CartCreate) error
	DeleteCartById(cartId int64) error
	ClearUserCart(userId int64) error
}

type CartService struct {
	cartRepository persistence.ICartRepository
}

func NewCartService(cartRepository persistence.ICartRepository) ICartService {
	return &CartService{
		cartRepository: cartRepository,
	}
}

func (cartService *CartService) GetCartById(cartId int64) domain.Cart {
	return cartService.cartRepository.GetCartById(cartId)
}

func (cartService *CartService) CreateCart(cart model.CartCreate) error {

	if validationErr := validation.ValidateCartCreate(cart); validationErr != nil {
		return validationErr
	}

	err := cartService.cartRepository.CreateCart(domain.Cart{
		UserId:    cart.UserId,
		CreatedAt: cart.CreatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (cartService *CartService) GetCartsByUserId(userId int64) []domain.Cart {
	return cartService.cartRepository.GetCartsByUserId(userId)
}

func (cartService *CartService) DeleteCartById(cartId int64) error {
	return cartService.cartRepository.DeleteCartById(cartId)
}

func (cartService *CartService) ClearUserCart(userId int64) error {
	return cartService.cartRepository.ClearUserCart(userId)
}
