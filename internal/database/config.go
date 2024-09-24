package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Raihanki/todolist/internal/config"
	_ "github.com/lib/pq"
)

func GetDatabaseConnection(cfg *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("could not open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("could not ping database: %v", err)
	}

	log.Println("connected to database")
	return db
}
