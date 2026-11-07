package lesson

import (
	"course-manager/internal/repository"
	"course-manager/internal/service"
)

var _ service.LessonService = (*Service)(nil)

type Service struct {
	lesson repository.LessonRepository
}

func New(lesson repository.LessonRepository) *Service {
	return &Service{lesson: lesson}
}
