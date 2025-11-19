package helper

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence/common"

	"github.com/jackc/pgx/v4"
)

func ScanProduct(row pgx.Row) (domain.Product, error) {
	var product domain.Product
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Slug,
		&product.Description,
		&product.Price,
		&product.BasePrice,
		&product.Discount,
		&product.ImageUrl,
		&product.MetaDescription,
		&product.StockQuantity,
		&product.IsActive,
		&product.IsFeatured,
		&product.CategoryId,
		&product.StoreId,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.Product{}, common.ErrProductNotFound
		}
		return product, common.WrapError("scan product", err)
	}
	return product, nil
}

func ScanStore(row pgx.Row) (domain.Store, error) {
	var store domain.Store
	err := row.Scan(
		&store.Id,
		&store.Name,
		&store.Slug,
		&store.Description,
		&store.LogoUrl,
		&store.ContactEmail,
		&store.ContactPhone,
		&store.ContactAddress,
		&store.IsActive,
		&store.CreatedAt,
		&store.UpdatedAt)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.Store{}, common.ErrStoreNotFound
		}
		return store, common.WrapError("scan store", err)
	}
	return store, nil
}

func ScanUser(row pgx.Row) (domain.User, error) {
	var user domain.User
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.User{}, common.ErrUserNotFound
		}
		return user, common.WrapError("scan user", err)
	}
	return user, nil
}

func ScanCart(row pgx.Row) (domain.Cart, error) {
	var cart domain.Cart
	err := row.Scan(&cart.Id, &cart.UserId, &cart.CreatedAt)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.Cart{}, common.ErrCartNotFound
		}
		return cart, common.WrapError("scan cart", err)
	}
	return cart, nil
}

func ScanCartItem(row pgx.Row) (domain.CartItem, error) {
	var cartItem domain.CartItem
	err := row.Scan(&cartItem.Id, &cartItem.CartId, &cartItem.ProductId, &cartItem.Quantity)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.CartItem{}, common.ErrCartItemNotFound
		}
		return cartItem, common.WrapError("scan cart item", err)
	}
	return cartItem, nil
}

func ScanOrder(row pgx.Row) (domain.Order, error) {
	var order domain.Order
	err := row.Scan(&order.Id, &order.UserId, &order.TotalPrice, &order.Status, &order.CreatedAt)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.Order{}, common.ErrOrderNotFound
		}
		return order, common.WrapError("scan order", err)
	}
	return order, nil
}

func ScanOrderItem(row pgx.Row) (domain.OrderItem, error) {
	var orderItem domain.OrderItem
	err := row.Scan(&orderItem.Id, &orderItem.OrderId, &orderItem.ProductId, &orderItem.Quantity, &orderItem.Price)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.OrderItem{}, common.ErrOrderItemNotFound
		}
		return orderItem, common.WrapError("scan order item", err)
	}
	return orderItem, nil
}

func ScanCategory(row pgx.Row) (domain.Category, error) {
	var category domain.Category
	err := row.Scan(&category.Id, &category.Name, &category.Description, &category.IsActive)
	if err != nil {
		if err.Error() == common.NOT_FOUND {
			return domain.Category{}, common.ErrCategoryNotFound
		}
		return category, common.WrapError("scan category", err)
	}
	return category, nil
}
