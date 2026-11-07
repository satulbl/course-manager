package lesson

import (
	"context"
	"course-manager/internal/models"
	"fmt"
)

const maxRowsLimit = 100

func (s *Service) List(ctx context.Context, limit, offset int) ([]models.Lesson, int64, error) {
	if limit < 0 || limit > maxRowsLimit {
		return nil, 0, fmt.Errorf("%w: limit must be at most %d", models.ErrInvalidInput, maxRowsLimit)
	}
	if offset < 0 {
		return nil, 0, fmt.Errorf("%w: offset must be zero or greater", models.ErrInvalidInput)
	}

	lessons, err := s.lesson.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list lessons: %w", err)
	}

	total, err := s.lesson.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return lessons, total, nil
}
