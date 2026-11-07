package repository

import (
	"context"
	"course-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

type Repositories struct {
	User    UserRepository
	Mentor  MentorRepository
	Session SessionRepository
	Token   TokenRepository
	Course  CourseRepository
	Lesson  LessonRepository
}

type UserRepository interface {
	Count(ctx context.Context) (int64, error)
	Create(ctx context.Context, u *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	List(ctx context.Context, limit, offset int) ([]models.User, error)
}

type MentorRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type SessionRepository interface {
	Create(ctx context.Context, s *models.Session) error
	List(ctx context.Context, userID uuid.UUID) ([]models.Session, error)
	Extend(ctx context.Context, id string, expiresAt time.Time) error
	Delete(ctx context.Context, userID uuid.UUID, id string) error
	DeleteOthers(ctx context.Context, userID uuid.UUID, keepID string) ([]string, error)
	DeleteExpired(ctx context.Context) (int64, error)
}

type TokenRepository interface {
	SavePair(ctx context.Context, sessionID string, pair *models.TokenPair, accessTTL, refreshTTL time.Duration) error
	IsAccessValid(ctx context.Context, sessionID, accessToken string) (bool, error)
	UseRefresh(ctx context.Context, sessionID, refreshToken string) (bool, error)
	Delete(ctx context.Context, sessionID string) error
	DeleteMany(ctx context.Context, sessionIDs []string) error
}

type CourseRepository interface {
	Count(ctx context.Context) (int64, error)
	List(ctx context.Context, limit, offset int) ([]models.Course, error)
	Create(ctx context.Context, c *models.Course) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Course, error)
	Update(ctx context.Context, c *models.Course) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type LessonRepository interface {
	Create(ctx context.Context, l *models.Lesson) error
	GetById(ctx context.Context, id uuid.UUID) (*models.Lesson, error)
	List(ctx context.Context, limit, offset int) ([]models.Lesson, error)
	Count(ctx context.Context) (int64, error)
}
