package dto

type CreateChecklistItemRequest struct {
	ChecklistId int `json:"checklist_id"`
}

type UpdateChecklistItemRequest struct {
	ItemName string `json:"item_name"`
}

type ChecklistItemResponse struct {
	ID          int    `json:"id"`
	ChecklistId int    `json:"checklist_id"`
	ItemName    string `json:"item_name"`
	IsCompleted bool   `json:"is_completed"`
}
