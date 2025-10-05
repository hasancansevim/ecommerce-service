package helper

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/common"

	"github.com/jackc/pgx/v4"
)

func ScanProduct(row pgx.Row) (domain.Product, error) {
	var product domain.Product
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Discount, &product.Store)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.Product{}, common.ErrProductNotFound
		}
		return product, common.WrapError("scan product", err)
	}
	return product, nil
}
