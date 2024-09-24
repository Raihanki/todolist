package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Raihanki/todolist/internal/config"
	"github.com/Raihanki/todolist/internal/database"
)

type ApplicationConfig struct {
	DB *sql.DB
}

func main() {
	db := database.GetDatabaseConnection(config.Get())
	defer db.Close()
	appCfg := &ApplicationConfig{DB: db}

	server := &http.Server{
		Addr:    ":" + config.Get().Server.Port,
		Handler: appCfg.Routes(),
	}

	log.Println("server started on port", config.Get().Server.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
