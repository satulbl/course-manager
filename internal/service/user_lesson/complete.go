package userlessons

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Complete(ctx context.Context, user models.CurrentUser, id uuid.UUID) error {
	if id == uuid.Nil {
		return fmt.Errorf("%w: id is required", models.ErrInvalidInput)
	}

	ul, err := s.userLessons.GetById(ctx, id)
	if err != nil {
		return err
	}

	if !user.CanCompleteLesson(ul.User.ID) {
		return models.ErrForbidden
	}

	if ul.IsCompleted() {
		return models.ErrAlreadyCompleted
	}

	return s.userLessons.Complete(ctx, id)
}
