package usercourse

import (
	"course-manager/internal/repository"
	"course-manager/internal/service"
)

var _ service.UserCoursesService = (*Service)(nil)

type Service struct {
	userCourses repository.UserCourseRepository
}

func New(userCourses repository.UserCourseRepository) *Service {
	return &Service{userCourses: userCourses}
}

