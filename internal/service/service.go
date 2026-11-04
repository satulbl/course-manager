package service

import (
	"context"
	"course-manager/internal/models"

	"github.com/google/uuid"
)

type Services struct {
	Auth   AuthService
	Course CourseService
	User   UserService
}

type AuthService interface {
	Register(ctx context.Context, in models.RegisterInput) (*models.User, error)
	Login(ctx context.Context, in models.LoginInput) (*models.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (*models.TokenPair, error)
	Authenticate(ctx context.Context, accessToken string) (models.CurrentUser, error)
}

type CourseService interface {
	Create(ctx context.Context, user models.CurrentUser, in models.CourseInput) error
	List(ctx context.Context, limit, offset int) ([]models.Course, int64, error)
	Update(ctx context.Context, user models.CurrentUser, id uuid.UUID, in models.CourseInput) error
	Delete(ctx context.Context, user models.CurrentUser, id uuid.UUID) error
}

type UserService interface {
	GetById(ctx context.Context, id uuid.UUID) (*models.User, error)
	List(ctx context.Context, limit, offset int) ([]models.User, int64, error)
}
