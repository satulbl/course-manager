package usercourse

import (
	"context"
	"course-manager/internal/models"

	"github.com/google/uuid"
)

func (s *Service) Approve(ctx context.Context, user models.CurrentUser, id uuid.UUID) error {
	if !user.IsAdmin() {
		return models.ErrForbidden
	}

	if err := s.userCourses.Approve(ctx, id); err != nil {
		return err
	}

	return nil
}
