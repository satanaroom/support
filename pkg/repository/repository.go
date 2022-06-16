package repository

import (
	"github.com/jmoiron/sqlx"
	"support"
)

type SupportItem interface {
	Create(item support.Support) (int, error)
	GetAll() ([]support.Support, error)
	GetItemById(itemId int) (support.Support, error)
	Update(id int, input support.UpdateSupportItem) error
}

type Repository struct {
	SupportItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SupportItem: NewSupportItemPostgres(db),
	}
}
