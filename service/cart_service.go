package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/internal/rules"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
	"time"
)

type ICartService interface {
	GetCartsByUserId(userId int64) []dto.CartResponse
	GetCartById(cartId int64) dto.CartResponse
	CreateCart(cart dto.CreateCartRequest) (dto.CartResponse, error)
	DeleteCartById(cartId int64) error
	ClearUserCart(userId int64) error
}

type CartService struct {
	cartRepository persistence.ICartRepository
	validator      *rules.CartRules
}

func NewCartService(cartRepository persistence.ICartRepository) ICartService {
	return &CartService{
		cartRepository: cartRepository,
		validator:      rules.NewCartRules(),
	}
}

func (cartService *CartService) GetCartById(cartId int64) dto.CartResponse {
	cart := cartService.cartRepository.GetCartById(cartId)
	cartDto := dto.CartResponse{
		UserId:    cart.UserId,
		CreatedAt: cart.CreatedAt,
	}
	return cartDto
}

func (cartService *CartService) CreateCart(cart dto.CreateCartRequest) (dto.CartResponse, error) {
	if validationErr := cartService.validator.ValidateStructure(cart); validationErr != nil {
		return dto.CartResponse{}, _errors.NewBadRequest(validationErr.Error())
	}

	createdCart, err := cartService.cartRepository.CreateCart(domain.Cart{
		UserId:    cart.UserId,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return dto.CartResponse{}, _errors.NewBadRequest(err.Error())
	}

	createdCartDto := dto.CartResponse{
		UserId:    createdCart.UserId,
		CreatedAt: createdCart.CreatedAt,
	}

	return createdCartDto, nil
}

func (cartService *CartService) GetCartsByUserId(userId int64) []dto.CartResponse {
	carts := cartService.cartRepository.GetCartsByUserId(userId)

	cartsDto := make([]dto.CartResponse, 0, len(carts))
	for _, cart := range carts {
		cartsDto = append(cartsDto, dto.CartResponse{
			UserId:    cart.UserId,
			CreatedAt: cart.CreatedAt,
		})
	}
	return cartsDto
}

func (cartService *CartService) DeleteCartById(cartId int64) error {
	return cartService.cartRepository.DeleteCartById(cartId)
}

func (cartService *CartService) ClearUserCart(userId int64) error {
	return cartService.cartRepository.ClearUserCart(userId)
}
