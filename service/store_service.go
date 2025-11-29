package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/persistence"
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
}

func NewStoreService(storeRepository persistence.IStoreRepository) IStoreService {
	return &StoreService{
		storeRepository: storeRepository,
	}
}

func (s *StoreService) GetAllStores() []dto.StoreResponse {
	stores := s.storeRepository.GetAllStores()
	storesDto := make([]dto.StoreResponse, 0, len(stores))

	for _, store := range stores {
		storesDto = append(storesDto, dto.StoreResponse{
			Name:           store.Name,
			Slug:           store.Slug,
			Description:    store.Description,
			IsActive:       store.IsActive,
			ContactEmail:   store.ContactEmail,
			ContactPhone:   store.ContactPhone,
			ContactAddress: store.ContactAddress,
			LogoUrl:        store.LogoUrl,
			UpdatedAt:      store.UpdatedAt,
			CreatedAt:      store.CreatedAt,
		})
	}
	return storesDto
}
func (s *StoreService) GetStoreById(storeId uint) (dto.StoreResponse, error) {
	store, err := s.storeRepository.GetStoreById(storeId)
	if err != nil {
		return dto.StoreResponse{}, err
	}
	storeDto := dto.StoreResponse{
		Name:           store.Name,
		Slug:           store.Slug,
		Description:    store.Description,
		IsActive:       store.IsActive,
		ContactEmail:   store.ContactEmail,
		ContactPhone:   store.ContactPhone,
		ContactAddress: store.ContactAddress,
		LogoUrl:        store.LogoUrl,
		UpdatedAt:      store.UpdatedAt,
		CreatedAt:      store.CreatedAt,
	}
	return storeDto, nil
}
func (s *StoreService) AddStore(store dto.CreateStoreRequest) (dto.StoreResponse, error) {
	addedStore, err := s.storeRepository.AddStore(domain.Store{
		Name:           store.Name,
		Slug:           store.Slug,
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
		return dto.StoreResponse{}, err
	}
	addedStoreDto := dto.StoreResponse{
		Name:           addedStore.Name,
		Slug:           addedStore.Slug,
		Description:    addedStore.Description,
		IsActive:       addedStore.IsActive,
		ContactEmail:   addedStore.ContactEmail,
		ContactPhone:   addedStore.ContactPhone,
		ContactAddress: addedStore.ContactAddress,
		LogoUrl:        addedStore.LogoUrl,
		CreatedAt:      addedStore.CreatedAt,
		UpdatedAt:      addedStore.UpdatedAt,
	}
	return addedStoreDto, nil
}
func (s *StoreService) DeleteStoreById(storeId uint) error {
	return s.storeRepository.DeleteStoreById(storeId)
}
func (s *StoreService) UpdateStoreById(id uint, store dto.CreateStoreRequest) (dto.StoreResponse, error) {
	updatedStore, err := s.storeRepository.UpdateStoreById(id, domain.Store{
		Name:           store.Name,
		Slug:           store.Slug,
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
		return dto.StoreResponse{}, err
	}
	updatedStoreDto := dto.StoreResponse{
		Name:           updatedStore.Name,
		Slug:           updatedStore.Slug,
		Description:    updatedStore.Description,
		IsActive:       updatedStore.IsActive,
		ContactEmail:   updatedStore.ContactEmail,
		ContactPhone:   updatedStore.ContactPhone,
		ContactAddress: updatedStore.ContactAddress,
		LogoUrl:        updatedStore.LogoUrl,
		CreatedAt:      updatedStore.CreatedAt,
		UpdatedAt:      updatedStore.UpdatedAt,
	}
	return updatedStoreDto, nil
}
