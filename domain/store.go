package domain

import "time"

type Store struct {
	Id             uint
	Name           string
	Slug           string
	Description    string
	LogoUrl        string
	ContactEmail   string
	ContactPhone   string
	ContactAddress string
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
