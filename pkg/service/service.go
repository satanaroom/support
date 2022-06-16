package service

import (
	"support"
	"support/pkg/repository"
)

type SupportItem interface {
	Create(item support.Support) (int, error)
	GetAll() ([]support.Support, error)
	GetItemById(itemId int) (support.Support, error)
	Update(id int, input support.UpdateSupportItem) error
}

type Service struct {
	SupportItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SupportItem: NewSupportItemService(repos.SupportItem),
	}
}
