package userlessons

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, userId uuid.UUID, in models.UserLessonInput) (uuid.UUID, error) {
	if userId == uuid.Nil {
		return uuid.Nil, fmt.Errorf("%w: user id is required", models.ErrInvalidInput)
	}

	if err := in.Validate(); err != nil {
		return uuid.Nil, err
	}

	ul := models.UserLesson{
		ID:     uuid.New(),
		User:   models.ApplicationUser{ID: userId},
		Lesson: models.ApplicationLesson{ID: in.LessonID},
	}

	err := s.userLessons.Create(ctx, &ul)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create user lesson: %w", err)
	}

	return ul.ID, nil
}
