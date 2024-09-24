package dto

import "time"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
