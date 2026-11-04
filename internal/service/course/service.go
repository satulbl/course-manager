package course

import (
	"context"
	"course-manager/internal/models"
	"course-manager/internal/repository"
	"course-manager/internal/service"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var _ service.CourseService = (*Service)(nil)

type Service struct {
	courses repository.CourseRepository
	mentor  repository.MentorRepository
}

func New(courses repository.CourseRepository, mentor repository.MentorRepository) *Service {
	return &Service{
		courses: courses,
		mentor:  mentor,
	}
}

func (s *Service) getCourseForEdit(ctx context.Context, user models.CurrentUser, id uuid.UUID) (*models.Course, error) {
	course, err := s.courses.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !user.CanEditCourse(course.AuthorID) {
		return nil, models.ErrForbidden
	}
	return course, nil
}

func (s *Service) mentorExist(ctx context.Context, id uuid.UUID) error {
	_, err := s.mentor.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("%w: mentor not found", models.ErrInvalidInput)
		}
		return fmt.Errorf("get mentor: %w", err)
	}

	return nil
}
