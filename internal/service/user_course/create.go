package usercourse

import (
	"context"
	"course-manager/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, in models.UserCourseInput) (uuid.UUID, error) {
	if err := in.Validate(); err != nil {
		return uuid.Nil, err
	}

	uc := models.UserCourse{
		ID: uuid.New(),
		User: models.ApplicationUser{
			ID: in.UserID,
		},
		Course: models.ApplicationCourse{
			ID: in.CourseID,
		},
	}

	if err := s.userCourses.Create(ctx, &uc); err != nil {
		return uuid.Nil, fmt.Errorf("create user course: %w", err)
	}

	return uc.ID, nil
}