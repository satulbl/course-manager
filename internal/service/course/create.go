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
		MentorId:    in.MentorId,
		StartDate:   in.StartDate,
		EndDate:     in.EndDate,
	}

	if err := s.mentorExist(ctx, in.MentorId); err != nil {
		return err
	}

	if err := s.courses.Create(ctx, course); err != nil {
		return fmt.Errorf("create course: %w", err)
	}

	return nil
}
