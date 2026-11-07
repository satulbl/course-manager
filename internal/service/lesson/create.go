package lesson

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, user models.CurrentUser, in models.LessonInput) (uuid.UUID, error) {
	if !user.CanCreate() {
		return uuid.Nil, models.ErrForbidden
	}

	if err := in.Validate(); err != nil {
		return uuid.Nil, err
	}

	lesson := &models.Lesson{
		ID:           uuid.New(),
		Title:        in.Title,
		Description:  in.Description,
		Content:      in.Content,
		LessonNumber: in.LessonNumber,
		Course: models.LessonCourse{
			ID: in.CourseID,
		},
	}

	if err := s.lesson.Create(ctx, lesson); err != nil {
		return uuid.Nil, fmt.Errorf("create lesson: %w", err)
	}

	return lesson.ID, nil
}
