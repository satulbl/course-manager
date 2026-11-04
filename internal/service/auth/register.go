package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"course-manager/internal/models"
)

func (s *Service) Register(ctx context.Context, in models.RegisterInput) (*models.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	u := &models.User{
		ID:           uuid.New(),
		Email:        in.Email,
		PasswordHash: string(hash),
		Name:         in.Name,
	}
	// Занятость email не проверяем заранее: это сделает UNIQUE в БД (models.ErrEmailTaken).
	if err := s.users.Create(ctx, u); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return u, nil
}
