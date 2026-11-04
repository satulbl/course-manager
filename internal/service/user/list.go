package user

import (
	"context"
	"course-manager/internal/models"
	"fmt"
)

const maxRowsLimit = 100

func (s *Service) List(ctx context.Context, limit, offset int) ([]models.User, int64, error) {
	if limit < 1 || limit > maxRowsLimit {
		return nil, 0, fmt.Errorf("%w: limit must be between 1 and %d", models.ErrInvalidInput, maxRowsLimit)
	}
	if offset < 0 {
		return nil, 0, fmt.Errorf("%w: offset must be zero or greater", models.ErrInvalidInput)
	}

	users, err := s.users.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get users: %w", err)
	}

	count, err := s.users.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
