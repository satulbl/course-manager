package handler

import (
	"net/http"

	"course-manager/internal/handler/middleware"
	v1 "course-manager/internal/handler/v1"
	"course-manager/internal/service"
)

func NewRouter(services *service.Services) http.Handler {
	api := v1.NewHandler(services)
	requireAuth := middleware.Auth(services.Auth)
	protected := func(h http.HandlerFunc) http.Handler {
		return requireAuth(h)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/auth/register", api.Register)
	mux.HandleFunc("POST /api/v1/auth/login", api.Login)
	// Refresh без requireAuth: в заголовке у него refresh-токен, а не access.
	mux.HandleFunc("POST /api/v1/auth/refresh", api.Refresh)

	mux.Handle("GET /api/v1/users-list", protected(api.ListUsers))

	mux.Handle("GET /api/v1/courses", protected(api.ListCourses))
	mux.Handle("POST /api/v1/courses", protected(api.CreateCourse))
	mux.Handle("PUT /api/v1/courses/{id}", protected(api.UpdateCourse))
	mux.Handle("DELETE /api/v1/courses/{id}", protected(api.DeleteCourse))

	mux.Handle("POST /api/v1/lessons", protected(api.CreateLesson))
	mux.Handle("GET /api/v1/lessons/{id}", protected(api.GetLessonById))
	mux.Handle("GET /api/v1/lessons", protected(api.ListLessons))

	// Logging снаружи: он видит и ответ 500, который отдаст Recover после паники.
	return middleware.Logging(middleware.Recover(mux))
}
