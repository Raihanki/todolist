package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Raihanki/todolist/internal/dto"
	"github.com/Raihanki/todolist/internal/helpers"
	"github.com/Raihanki/todolist/internal/services"
)

type UserHandler struct {
	UserService services.UserService
}

func (service *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	request := dto.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response, err := service.UserService.RegisterUser(request)
	if err != nil {
		helpers.JsonResponse(w, http.StatusInternalServerError, "", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, http.StatusText(http.StatusOK), response)
}

func (service *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	request := dto.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response, err := service.UserService.LoginUser(request)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) || errors.Is(err, errors.New("invalid password")) {
			helpers.JsonResponse(w, http.StatusBadRequest, "invalid username or password", nil)
			return
		}
		helpers.JsonResponse(w, http.StatusInternalServerError, "", nil)
		return
	}

	helpers.JsonResponse(w, http.StatusOK, http.StatusText(http.StatusOK), response)
}
