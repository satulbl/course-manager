package course

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Delete(ctx context.Context, user models.CurrentUser, id uuid.UUID) error {
	_, err := s.getCourseForEdit(ctx, user, id)
	if err != nil {
		return fmt.Errorf("get course: %w", err)
	}

	err = s.courses.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("delete course: %w", err)
	}

	return nil
}
