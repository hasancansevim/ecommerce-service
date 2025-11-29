package dto

type CategoryResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
