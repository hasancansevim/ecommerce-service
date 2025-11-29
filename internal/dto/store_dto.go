package dto

import "time"

type StoreResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	Description    string    `json:"description"`
	LogoUrl        string    `json:"logo_url"`
	ContactEmail   string    `json:"contact_email"`
	ContactPhone   string    `json:"contact_phone"`
	ContactAddress string    `json:"contact_address"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateStoreRequest struct {
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	Description    string `json:"description"`
	LogoUrl        string `json:"logo_url"`
	ContactEmail   string `json:"contact_email"`
	ContactPhone   string `json:"contact_phone"`
	ContactAddress string `json:"contact_address"`
	IsActive       bool   `json:"is_active"`
}
