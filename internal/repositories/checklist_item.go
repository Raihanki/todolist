package repositories

import (
	"database/sql"

	"github.com/Raihanki/todolist/internal/domain"
)

type ChecklistItemRepository interface {
	CreateChecklistItem(checklist domain.ChecklistItem) error
}

type ChecklistItemRepositoryImpl struct {
	DB *sql.DB
}

func NewChecklistItemRepository(db *sql.DB) *ChecklistItemRepositoryImpl {
	return &ChecklistItemRepositoryImpl{DB: db}
}

func (r *ChecklistItemRepositoryImpl) CreateChecklistItem(checklist domain.ChecklistItem) error {
	query := "INSERT INTO checklist_items (checklist_id, name) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, checklist.ChecklistId, checklist.ItemName)
	if err != nil {
		return err
	}

	return nil
}
