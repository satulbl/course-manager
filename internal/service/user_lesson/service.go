package userlessons

import (
	"course-manager/internal/repository"
	"course-manager/internal/service"
)

var _ service.UserLessonService = (*Service)(nil)

type Service struct {
	userLessons repository.UserLessonRepository
}

func New(userLessons repository.UserLessonRepository) *Service {
	return &Service{userLessons: userLessons}
}