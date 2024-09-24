package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Raihanki/todolist/internal/dto"
	"github.com/Raihanki/todolist/internal/helpers"
	"github.com/Raihanki/todolist/internal/services"
)

type ChecklistHandler struct {
	ChecklistService services.ChecklistService
}

func (service *ChecklistHandler) CreateChecklist(w http.ResponseWriter, r *http.Request, userId int) {
	var request dto.CreateChecklistRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = service.ChecklistService.CreateChecklist(request, userId)
	if err != nil {
		helpers.JsonResponse(w, http.StatusInternalServerError, "", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusCreated, http.StatusText(http.StatusCreated), nil)
}

func (service *ChecklistHandler) GetChecklist(w http.ResponseWriter, r *http.Request, userId int) {
	checklists, err := service.ChecklistService.GetChecklist(userId)
	if err != nil {
		helpers.JsonResponse(w, http.StatusInternalServerError, "", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, http.StatusText(http.StatusOK), checklists)
}

func (service *ChecklistHandler) DeleteChecklist(w http.ResponseWriter, r *http.Request, userId int) {
	cId := r.PathValue("checklist-id")
	checklistId, err := strconv.Atoi(cId)
	if err != nil {
		helpers.JsonResponse(w, http.StatusBadRequest, "invalid checklist id", nil)
		return
	}

	err = service.ChecklistService.DeleteChecklist(checklistId, userId)
	if err != nil {
		if errors.Is(err, errors.New("checklist not found")) {
			helpers.JsonResponse(w, http.StatusNotFound, "checklist not found", nil)
			return
		}
		helpers.JsonResponse(w, http.StatusInternalServerError, "", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, http.StatusText(http.StatusOK), nil)
}
