package course

import (
	"context"
	"course-manager/internal/models"
	"fmt"
)

const maxRowsLimit = 100

func (s *Service) List(ctx context.Context, limit, offset int) ([]models.Course, int64, error) {
	if limit < 1 || limit > maxRowsLimit {
		return nil, 0, fmt.Errorf("%w: limit must be between 1 and %d", models.ErrInvalidInput, maxRowsLimit)
	}
	if offset < 0 {
		return nil, 0, fmt.Errorf("%w: offset must be zero or greater", models.ErrInvalidInput)
	}

	courses, err := s.courses.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("list course: %w", err)
	}

	total, err := s.courses.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return courses, total, nil
}
