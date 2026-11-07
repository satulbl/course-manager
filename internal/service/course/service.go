package course

import (
	"context"
	"course-manager/internal/models"
	"course-manager/internal/repository"
	"course-manager/internal/service"

	"github.com/google/uuid"
)

var _ service.CourseService = (*Service)(nil)

type Service struct {
	courses repository.CourseRepository
}

func New(courses repository.CourseRepository) *Service {
	return &Service{courses: courses}
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