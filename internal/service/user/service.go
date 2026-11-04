package user

import "course-manager/internal/repository"

type Service struct {
	users repository.UserRepository
}

func New(users repository.UserRepository) *Service {
	return &Service{users: users}
}
