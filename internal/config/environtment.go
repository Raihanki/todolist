package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   Server
	Database Database
	App      App
	JWT      JWT
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWT struct {
	Secret string
}

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env file: %v", err)
	}

	return &Config{
		App: App{
			Name: os.Getenv("APP_NAME"),
		},
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		JWT: JWT{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}
