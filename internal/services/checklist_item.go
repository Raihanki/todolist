package services

import (
	"github.com/Raihanki/todolist/internal/domain"
	"github.com/Raihanki/todolist/internal/dto"
	"github.com/Raihanki/todolist/internal/repositories"
)

type ChecklistItemService interface {
	CreateChecklistItem(request dto.CreateChecklistRequest, userId int) error
}

type ChecklistItemServiceImpl struct {
	repo repositories.ChecklistItemRepository
}

func NewChecklistItemService(repo repositories.ChecklistItemRepository) *ChecklistItemServiceImpl {
	return &ChecklistItemServiceImpl{repo: repo}
}

func (s *ChecklistItemServiceImpl) CreateChecklistItem(request dto.CreateChecklistRequest, userId int) error {
	checklistItem := domain.ChecklistItem{
		ItemName: request.ItemName,
	}
	err := s.repo.CreateChecklistItem(checklistItem)
	if err != nil {
		return err
	}

	return nil
}
