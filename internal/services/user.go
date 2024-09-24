package services

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Raihanki/todolist/internal/domain"
	"github.com/Raihanki/todolist/internal/dto"
	"github.com/Raihanki/todolist/internal/helpers"
	"github.com/Raihanki/todolist/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(request dto.RegisterRequest) (dto.AuthResponse, error)
	LoginUser(request dto.LoginRequest) (dto.AuthResponse, error)
}

type UserServiceImpl struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) RegisterUser(request dto.RegisterRequest) (dto.AuthResponse, error) {
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password ERR:%v", err)
		return dto.AuthResponse{}, errors.New("failed to hash password")
	}

	//create user
	user := domain.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Email:    request.Email,
	}
	err = s.repo.CreateUser(user)
	if err != nil {
		log.Printf("failed to create user ERR:%v", err)
		return dto.AuthResponse{}, errors.New("failed to create user")
	}

	//generate token
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		log.Printf("failed to generate token ERR:%v", err)
		return dto.AuthResponse{}, errors.New("failed to generate token")
	}

	return dto.AuthResponse{
		Token:     token,
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}, nil
}

func (s *UserServiceImpl) LoginUser(request dto.LoginRequest) (dto.AuthResponse, error) {
	//get user by username
	user, err := s.repo.GetUserByUsername(request.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.AuthResponse{}, errors.New("user not found")
		}
		log.Printf("failed to get user ERR:%v", err)
		return dto.AuthResponse{}, errors.New("failed to get user")
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		log.Printf("invalid password ERR:%v", err)
		return dto.AuthResponse{}, errors.New("invalid password")
	}

	//generate token
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		log.Printf("failed to generate token ERR:%v", err)
		return dto.AuthResponse{}, errors.New("failed to generate token")
	}

	return dto.AuthResponse{
		Token:     token,
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}, nil
}
