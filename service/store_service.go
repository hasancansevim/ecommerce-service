package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
)

type IStoreService interface {
	GetAllStores() []domain.Store
	GetStoreById(storeId uint) (domain.Store, error)
	AddStore(store model.StoreCreate) error
	DeleteStoreById(storeId uint) error
	UpdateStoreById(id uint, store model.StoreCreate) error
}

type StoreService struct {
	storeRepository persistence.IStoreRepository
}

func NewStoreService(storeRepository persistence.IStoreRepository) IStoreService {
	return &StoreService{
		storeRepository: storeRepository,
	}
}

func (s *StoreService) GetAllStores() []domain.Store {
	return s.storeRepository.GetAllStores()
}
func (s *StoreService) GetStoreById(storeId uint) (domain.Store, error) {
	return s.storeRepository.GetStoreById(storeId)
}
func (s *StoreService) AddStore(store model.StoreCreate) error {
	err := s.storeRepository.AddStore(domain.Store{
		Name:           store.Name,
		Slug:           store.Slug,
		LogoUrl:        store.LogoUrl,
		ContactAddress: store.ContactAddress,
		ContactEmail:   store.ContactEmail,
		ContactPhone:   store.ContactPhone,
		IsActive:       store.IsActive,
		Description:    store.Description,
		CreatedAt:      store.CreatedAt,
		UpdatedAt:      store.UpdatedAt,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *StoreService) DeleteStoreById(storeId uint) error {
	return s.storeRepository.DeleteStoreById(storeId)
}
func (s *StoreService) UpdateStoreById(id uint, store model.StoreCreate) error {
	return s.storeRepository.UpdateStoreById(id, domain.Store{
		Name:           store.Name,
		Slug:           store.Slug,
		LogoUrl:        store.LogoUrl,
		ContactAddress: store.ContactAddress,
		ContactEmail:   store.ContactEmail,
		ContactPhone:   store.ContactPhone,
		IsActive:       store.IsActive,
		Description:    store.Description,
		CreatedAt:      store.CreatedAt,
		UpdatedAt:      store.UpdatedAt,
	})
}
