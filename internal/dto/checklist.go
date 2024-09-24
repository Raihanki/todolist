package dto

type CreateChecklistRequest struct {
	ItemName string `json:"item_name"`
}

type ChecklistResponse struct {
	ID       int    `json:"id"`
	UserId   int    `json:"user_id"`
	ItemName string `json:"item_name"`
}
