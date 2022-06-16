package service

import (
	"support"
	"support/pkg/repository"
)

type SupportService struct {
	repo repository.SupportItem
}

func NewSupportItemService(repo repository.SupportItem) *SupportService {
	return &SupportService{repo: repo}
}

func (s *SupportService) Create(item support.Support) (int, error) {
	return s.repo.Create(item)
}

func (s *SupportService) GetAll() ([]support.Support, error) {
	return s.repo.GetAll()
}

func (s *SupportService) GetItemById(itemId int) (support.Support, error) {
	return s.repo.GetItemById(itemId)
}

func (s *SupportService) Update(id int, input support.UpdateSupportItem) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
