package repositories

import (
	"database/sql"

	"github.com/Raihanki/todolist/internal/domain"
)

type ChecklistRepository interface {
	CreateChecklist(checklist domain.Checklist, userId int) error
	GetChecklist(userId int) ([]domain.Checklist, error)
	DeteleChecklist(id int, userId int) error
}

type ChecklistRepositoryImpl struct {
	DB *sql.DB
}

func NewChecklistRepository(db *sql.DB) *ChecklistRepositoryImpl {
	return &ChecklistRepositoryImpl{DB: db}
}

func (r *ChecklistRepositoryImpl) CreateChecklist(checklist domain.Checklist, userId int) error {
	query := "INSERT INTO checklists (name, user_id) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, checklist.Title, userId)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChecklistRepositoryImpl) GetChecklist(userId int) ([]domain.Checklist, error) {
	var checklists []domain.Checklist
	query := "SELECT id, title, user_id FROM checklists WHERE user_id = $1"
	rows, err := r.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var checklist domain.Checklist
		err := rows.Scan(&checklist.ID, &checklist.Title, &checklist.UserId)
		if err != nil {
			return nil, err
		}

		checklists = append(checklists, checklist)
	}

	return checklists, nil
}

func (r *ChecklistRepositoryImpl) DeteleChecklist(id int, userId int) error {
	query := "DELETE FROM checklists WHERE id = $1 AND user_id = $2"
	_, err := r.DB.Exec(query, id, userId)
	if err != nil {
		return err
	}

	return nil
}
