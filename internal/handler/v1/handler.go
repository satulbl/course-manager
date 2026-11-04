package v1

import "course-manager/internal/service"

// Каждая ручка делает три шага: разобрать запрос → вызвать сервис → ответить.
type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}
