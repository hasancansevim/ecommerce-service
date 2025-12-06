package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/internal/rules"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
	"go-ecommerce-service/pkg/util"
	"time"
)

type IStoreService interface {
	GetAllStores() []dto.StoreResponse
	GetStoreById(storeId uint) (dto.StoreResponse, error)
	AddStore(store dto.CreateStoreRequest) (dto.StoreResponse, error)
	DeleteStoreById(storeId uint) error
	UpdateStoreById(id uint, store dto.CreateStoreRequest) (dto.StoreResponse, error)
}

type StoreService struct {
	storeRepository persistence.IStoreRepository
	validator       *rules.StoreRules
}

func NewStoreService(storeRepository persistence.IStoreRepository) IStoreService {
	return &StoreService{
		storeRepository: storeRepository,
		validator:       rules.NewStoreRules(),
	}
}

func (s *StoreService) GetAllStores() []dto.StoreResponse {
	stores := s.storeRepository.GetAllStores()
	return convertToStoresResponse(stores)
}
func (s *StoreService) GetStoreById(storeId uint) (dto.StoreResponse, error) {
	store, err := s.storeRepository.GetStoreById(storeId)
	if err != nil {
		return dto.StoreResponse{}, _errors.NewInternalServerError(err)
	}

	return convertToStoreResponse(store), nil
}
func (s *StoreService) AddStore(store dto.CreateStoreRequest) (dto.StoreResponse, error) {
	if validationErr := s.validator.ValidateStructure(store); validationErr != nil {
		return dto.StoreResponse{}, _errors.NewBadRequest(validationErr.Error())
	}

	addedStore, err := s.storeRepository.AddStore(domain.Store{
		Name:           store.Name,
		Slug:           util.GenerateUniqueSlug(store.Name),
		LogoUrl:        store.LogoUrl,
		ContactAddress: store.ContactAddress,
		ContactEmail:   store.ContactEmail,
		ContactPhone:   store.ContactPhone,
		IsActive:       store.IsActive,
		Description:    store.Description,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return dto.StoreResponse{}, _errors.NewInternalServerError(err)
	}

	return convertToStoreResponse(addedStore), nil
}
func (s *StoreService) DeleteStoreById(storeId uint) error {
	return s.storeRepository.DeleteStoreById(storeId)
}
func (s *StoreService) UpdateStoreById(id uint, store dto.CreateStoreRequest) (dto.StoreResponse, error) {
	if validationErr := s.validator.ValidateStructure(store); validationErr != nil {
		return dto.StoreResponse{}, _errors.NewBadRequest(validationErr.Error())
	}

	updatedStore, err := s.storeRepository.UpdateStoreById(id, domain.Store{
		Name:           store.Name,
		Slug:           util.GenerateUniqueSlug(store.Name),
		LogoUrl:        store.LogoUrl,
		ContactAddress: store.ContactAddress,
		ContactEmail:   store.ContactEmail,
		ContactPhone:   store.ContactPhone,
		IsActive:       store.IsActive,
		Description:    store.Description,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	if err != nil {
		return dto.StoreResponse{}, _errors.NewInternalServerError(err)
	}

	return convertToStoreResponse(updatedStore), nil
}

func convertToStoreResponse(store domain.Store) dto.StoreResponse {
	return dto.StoreResponse{
		Id:          store.Id,
		Name:        store.Name,
		Slug:        store.Slug,
		Description: store.Description,
		IsActive:    store.IsActive,
	}
}

func convertToStoresResponse(stores []domain.Store) []dto.StoreResponse {
	{
		storesDto := make([]dto.StoreResponse, 0, len(stores))
		for _, store := range stores {
			storesDto = append(storesDto, convertToStoreResponse(store))
		}
		return storesDto
	}
}
