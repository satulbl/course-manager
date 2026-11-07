package lesson

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) GetById(ctx context.Context, id uuid.UUID) (*models.Lesson, error) {
	lesson, err := s.lesson.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get lesson by id: %w", err)
	}

	return lesson, nil
}
