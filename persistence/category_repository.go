package persistence

import (
	"context"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/helper"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ICategoryRepository interface {
	GetAllCategories() []domain.Category
	GetCategoryById(id int) (domain.Category, error)
	GetCategoriesByIsActive(isActive bool) ([]domain.Category, error)
	AddCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(categoryId uint, category domain.Category) (domain.Category, error)
	DeleteCategory(id uint) error
}

type CategoryRepository struct {
	dbPool  *pgxpool.Pool
	scanner *helper.GenericScanner[domain.Category]
}

func NewCategoryRepository(dbPool *pgxpool.Pool) ICategoryRepository {
	return &CategoryRepository{
		dbPool:  dbPool,
		scanner: helper.NewGenericScanner(dbPool, helper.ScanCategory),
	}
}

func (categoryRepository *CategoryRepository) GetAllCategories() []domain.Category {
	ctx := context.Background()

	categories, err := categoryRepository.scanner.QueryAndScan(ctx, "SELECT * FROM categories")

	if err != nil {
		return []domain.Category{}
	}
	return categories
}

func (categoryRepository *CategoryRepository) GetCategoryById(id int) (domain.Category, error) {
	ctx := context.Background()
	category, err := categoryRepository.scanner.QueryRowAndScan(ctx, "SELECT * FROM categories WHERE id = $1", id)
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (categoryRepository *CategoryRepository) GetCategoriesByIsActive(isActive bool) ([]domain.Category, error) {
	ctx := context.Background()
	categories, err := categoryRepository.scanner.QueryAndScan(ctx, "SELECT * FROM categories WHERE is_active = $1", isActive)
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}

func (categoryRepository *CategoryRepository) AddCategory(category domain.Category) (domain.Category, error) {
	ctx := context.Background()
	query := `INSERT INTO categories (name, description, is_active) 
              VALUES ($1, $2, $3) 
              RETURNING *`

	category, err := categoryRepository.scanner.QueryRowAndScan(ctx, query, category.Name, category.Description, category.IsActive)
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}
func (categoryRepository *CategoryRepository) UpdateCategory(categoryId uint, category domain.Category) (domain.Category, error) {
	ctx := context.Background()
	query := `UPDATE categories set name = $1, description = $2, is_active = $3 WHERE id = $4 RETURNING *`

	category, err := categoryRepository.scanner.QueryRowAndScan(ctx, query, category.Name, category.Description, category.IsActive, categoryId)
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}
func (categoryRepository *CategoryRepository) DeleteCategory(id uint) error {
	ctx := context.Background()
	query := `DELETE FROM categories WHERE id = $1`
	err := categoryRepository.scanner.ExecuteExec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
