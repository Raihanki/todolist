package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Raihanki/todolist/internal/dto"
	"github.com/Raihanki/todolist/internal/helpers"
	"github.com/Raihanki/todolist/internal/services"
)

type ChecklistItemHandler struct {
	ChecklistItemService services.ChecklistItemService
}

func NewChecklistItemHandler(service services.ChecklistItemService) *ChecklistItemHandler {
	return &ChecklistItemHandler{ChecklistItemService: service}
}

func (service *ChecklistItemHandler) CreateChecklistItem(w http.ResponseWriter, r *http.Request, userId int) {
	var request dto.CreateChecklistRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = service.ChecklistItemService.CreateChecklistItem(request, userId)
	if err != nil {
		helpers.JsonResponse(w, http.StatusInternalServerError, "", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusCreated, http.StatusText(http.StatusCreated), nil)
}
