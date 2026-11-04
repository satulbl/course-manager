package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	Id   int64
	Name string
	Code string
}

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	Role         UserRole
	CreatedAt    time.Time
}

func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
