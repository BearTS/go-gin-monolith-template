package models

import "github.com/golang-jwt/jwt/v4"

type AuthData struct {
	SessionPID string `json:"session_pid" binding:"required"`
	UserPID    string `json:"user_pid" binding:"required"`
	AdminPID   string `json:"admin_pid" binding:"required"`
	Sandbox    bool   `json:"sandbox" binding:"required"`
	Type       string `json:"type" binding:"required"`
	jwt.RegisteredClaims
}
