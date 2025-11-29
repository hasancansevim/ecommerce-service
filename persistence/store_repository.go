package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IStoreRepository interface {
	GetAllStores() []domain.Store
	GetStoreById(storeId uint) (domain.Store, error)
	AddStore(store domain.Store) (domain.Store, error)
	DeleteStoreById(storeId uint) error
	UpdateStoreById(id uint, store domain.Store) (domain.Store, error)
}

type StoreRepository struct {
	dbPool  *pgxpool.Pool
	scanner *helper.GenericScanner[domain.Store]
}

func NewStoreRepository(dbPool *pgxpool.Pool) IStoreRepository {
	return &StoreRepository{
		dbPool:  dbPool,
		scanner: helper.NewGenericScanner(dbPool, helper.ScanStore),
	}
}

func (storeRepository *StoreRepository) GetAllStores() []domain.Store {
	ctx := context.Background()

	stores, err := storeRepository.scanner.QueryAndScan(ctx, "SELECT * FROM stores")

	if err != nil {
		return []domain.Store{}
	}
	return stores
}
func (storeRepository *StoreRepository) GetStoreById(storeId uint) (domain.Store, error) {
	ctx := context.Background()
	query := `Select * from stores where id = $1`
	store, err := storeRepository.scanner.QueryRowAndScan(ctx, query, storeId)
	if err != nil {
		return domain.Store{}, err
	}
	return store, nil
}
func (storeRepository *StoreRepository) AddStore(store domain.Store) (domain.Store, error) {
	ctx := context.Background()
	query := `INSERT INTO stores (name,slug,description,logo_url,contact_email,contact_phone,contact_address,is_active,created_at,updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING *`
	store, err := storeRepository.scanner.QueryRowAndScan(ctx, query,
		store.Name, store.Slug, store.Description,
		store.LogoUrl, store.ContactEmail, store.ContactPhone,
		store.ContactAddress, store.IsActive,
		store.CreatedAt, store.UpdatedAt)
	if err != nil {
		return domain.Store{}, err
	}
	return store, nil
}
func (storeRepository *StoreRepository) DeleteStoreById(storeId uint) error {
	ctx := context.Background()
	query := `DELETE from stores where id = $1`
	err := storeRepository.scanner.ExecuteExec(ctx, query, storeId)
	if err != nil {
		return err
	}
	return nil
}
func (storeRepository *StoreRepository) UpdateStoreById(id uint, store domain.Store) (domain.Store, error) {
	ctx := context.Background()
	query := `UPDATE stores set name=$1,slug=$2, description=$3,logo_url=$4,contact_email=$5,contact_phone=$6,contact_address=$7,is_active=$8,created_at=$9 , updated_at =$10 WHERE id = $11`
	store, err := storeRepository.scanner.QueryRowAndScan(ctx, query,
		store.Name, store.Slug, store.Description,
		store.LogoUrl, store.ContactEmail, store.ContactPhone,
		store.ContactAddress, store.IsActive, store.CreatedAt, time.Now(), id)
	if err != nil {
		return domain.Store{}, err
	}
	return store, nil
}
