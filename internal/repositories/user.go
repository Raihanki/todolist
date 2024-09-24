package repositories

import (
	"database/sql"
	"log"

	"github.com/Raihanki/todolist/internal/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserByUsername(username string) (domain.User, error)
}

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) CreateUser(user domain.User) error {
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) returning id"
	var id int
	err := r.DB.QueryRow(query, user.Username, user.Password, user.Email).Scan(&id)
	if err != nil {
		log.Printf("failed to create user repo ERR:%v", err)
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserByUsername(username string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, username, password, email, created_at FROM users WHERE username = $1"
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
