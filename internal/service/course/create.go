package course

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, user models.CurrentUser, in models.CourseInput) error {
	if !user.CanCreate() {
		return models.ErrForbidden
	}
	if err := in.Validate(); err != nil {
		return err
	}

	course := &models.Course{
		ID:          uuid.New(),
		Title:       in.Title,
		Description: in.Description,
		Price:       in.Price,
		AuthorID:    user.ID,
		MentorID:    in.MentorID,
		StartDate:   in.StartDate,
		EndDate:     in.EndDate,
	}

	if err := s.courses.Create(ctx, course); err != nil {
		return fmt.Errorf("create course: %w", err)
	}
	return nil
}