package auth

import (
	"time"

	"course-manager/internal/models"
	"course-manager/internal/repository"
	"course-manager/internal/service"

	"github.com/google/uuid"
)

var _ service.AuthService = (*Service)(nil)

// TokenManager — подпись и разбор JWT. Интерфейс позволяет подставить фейк в тестах.
type TokenManager interface {
	GenerateAccess(userID uuid.UUID, roleID int64, sessionID string) (string, error)
	GenerateRefresh(userID uuid.UUID, sessionID string) (string, error)
	ParseAccess(token string) (models.CurrentUser, error)
	ParseRefresh(token string) (userID uuid.UUID, sessionID string, err error)
	AccessTTL() time.Duration
	RefreshTTL() time.Duration
}

type Service struct {
	users    repository.UserRepository
	sessions repository.SessionRepository
	tokens   repository.TokenRepository
	jwt      TokenManager
}

func New(
	users repository.UserRepository,
	sessions repository.SessionRepository,
	tokens repository.TokenRepository,
	jwt TokenManager,
) *Service {
	return &Service{
		users:    users,
		sessions: sessions,
		tokens:   tokens,
		jwt:      jwt,
	}
}
