package course

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Update(ctx context.Context, user models.CurrentUser, id uuid.UUID, in models.CourseInput) error {
	c, err := s.getCourseForEdit(ctx, user, id)
	if err != nil {
		return fmt.Errorf("get course: %w", err)
	}

	if c.MentorId != in.MentorId {
		if err := s.mentorExist(ctx, in.MentorId); err != nil {
			return err
		}
	}

	if err := c.Edit(in); err != nil {
		return err
	}

	if err := s.courses.Update(ctx, c); err != nil {
		return err
	}
	return nil
}
