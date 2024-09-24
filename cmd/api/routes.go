package main

import (
	"net/http"

	"github.com/Raihanki/todolist/cmd/api/handlers"
	middlewares "github.com/Raihanki/todolist/internal/middleware"
	"github.com/Raihanki/todolist/internal/repositories"
	"github.com/Raihanki/todolist/internal/services"
)

func (cfg *ApplicationConfig) Routes() http.Handler {
	mux := http.NewServeMux()

	// user routes
	userRepo := repositories.NewUserRepository(cfg.DB)
	userHandler := &handlers.UserHandler{
		UserService: services.NewUserService(userRepo),
	}
	mux.HandleFunc("POST /api/register", userHandler.RegisterUser)
	mux.HandleFunc("POST /api/login", userHandler.LoginUser)

	//checklist routes
	checklistRepo := repositories.NewChecklistRepository(cfg.DB)
	checklistHandler := &handlers.ChecklistHandler{
		ChecklistService: services.NewChecklistService(checklistRepo),
	}
	mux.HandleFunc("POST /api/checklist", middlewares.AuthenticateUsingToken(checklistHandler.CreateChecklist))
	mux.HandleFunc("GET /api/checklist", middlewares.AuthenticateUsingToken(checklistHandler.GetChecklist))
	mux.HandleFunc("DELETE /api/checklist/{checklist-id}", middlewares.AuthenticateUsingToken(checklistHandler.DeleteChecklist))

	//checklist item routes
	checklistItemRepo := repositories.NewChecklistItemRepository(cfg.DB)
	checklistItemHandler := &handlers.ChecklistItemHandler{
		ChecklistItemService: services.NewChecklistItemService(checklistItemRepo),
	}
	mux.HandleFunc("POST /api/checklist/{checklist-id}/item", middlewares.AuthenticateUsingToken(checklistItemHandler.CreateChecklistItem))

	return mux
}
