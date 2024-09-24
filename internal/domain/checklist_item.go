package domain

type ChecklistItem struct {
	ID          int
	ChecklistId int
	ItemName    string
	IsCompleted bool
}
