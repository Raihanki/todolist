package services

import (
	"database/sql"
	"errors"

	"github.com/Raihanki/todolist/internal/domain"
	"github.com/Raihanki/todolist/internal/dto"
	"github.com/Raihanki/todolist/internal/repositories"
)

type ChecklistService interface {
	CreateChecklist(request dto.CreateChecklistRequest, userId int) error
	GetChecklist(userId int) ([]dto.ChecklistResponse, error)
	DeleteChecklist(id int, userId int) error
}

type ChecklistServiceImpl struct {
	repo repositories.ChecklistRepository
}

func NewChecklistService(repo repositories.ChecklistRepository) *ChecklistServiceImpl {
	return &ChecklistServiceImpl{repo: repo}
}

func (s *ChecklistServiceImpl) CreateChecklist(request dto.CreateChecklistRequest, userId int) error {
	checklist := domain.Checklist{
		Title: request.Title,
	}
	err := s.repo.CreateChecklist(checklist, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *ChecklistServiceImpl) GetChecklist(userId int) ([]dto.ChecklistResponse, error) {
	checklists, err := s.repo.GetChecklist(userId)
	if err != nil {
		return nil, err
	}

	var response []dto.ChecklistResponse
	for _, checklist := range checklists {
		response = append(response, dto.ChecklistResponse{
			ID:     checklist.ID,
			UserId: checklist.UserId,
			Title:  checklist.Title,
		})
	}

	return response, nil
}

func (s *ChecklistServiceImpl) DeleteChecklist(id int, userId int) error {
	err := s.repo.DeteleChecklist(id, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("checklist not found")
		}
		return err
	}

	return nil
}
