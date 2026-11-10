package httpx

import (
	"course-manager/internal/models"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

// JSON отдаёт успешный ответ: данные кладутся в payload.
func JSON(w http.ResponseWriter, status int, payload any) {
	write(w, status, Response{Payload: payload})
}

func write(w http.ResponseWriter, status int, body Response) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("encode json response", "error", err)
	}
}

func Error(w http.ResponseWriter, r *http.Request, status int, message string, cause error) {
	level := slog.LevelWarn
	if status >= http.StatusInternalServerError {
		level = slog.LevelError
	}
	slog.Log(r.Context(), level, "request failed",
		"method", r.Method,
		"path", r.URL.Path,
		"status", status,
		"error", cause,
	)
	write(w, status, Response{Message: message})
}

func ServiceError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, models.ErrInvalidInput):
		Error(w, r, http.StatusBadRequest, err.Error(), err)
	case errors.Is(err, models.ErrInvalidCredentials):
		Error(w, r, http.StatusUnauthorized, models.ErrInvalidCredentials.Error(), err)
	case errors.Is(err, models.ErrInvalidToken):
		Error(w, r, http.StatusUnauthorized, models.ErrInvalidToken.Error(), err)
	case errors.Is(err, models.ErrEmailTaken):
		Error(w, r, http.StatusConflict, models.ErrEmailTaken.Error(), err)
	case errors.Is(err, models.ErrNotFound):
		Error(w, r, http.StatusNotFound, models.ErrNotFound.Error(), err)
	case errors.Is(err, models.ErrForbidden):
		Error(w, r, http.StatusForbidden, models.ErrForbidden.Error(), err)
	case errors.Is(err, models.ErrFileTooLarge):
		Error(w, r, http.StatusRequestEntityTooLarge, models.ErrFileTooLarge.Error(), err)
	case errors.Is(err, models.ErrUnsupportedFileType):
		Error(w, r, http.StatusUnsupportedMediaType, models.ErrUnsupportedFileType.Error(), err)
	case errors.Is(err, models.ErrFileNotAvailable):
		Error(w, r, http.StatusBadRequest, models.ErrFileNotAvailable.Error(), err)
	case errors.Is(err, models.ErrLessonNumberTaken):
		Error(w, r, http.StatusConflict, models.ErrLessonNumberTaken.Error(), err)
	case errors.Is(err, models.ErrCourseNotFound):
		Error(w, r, http.StatusBadRequest, models.ErrCourseNotFound.Error(), err)
	case errors.Is(err, models.ErrUserNotFound):
		Error(w, r, http.StatusBadRequest, models.ErrUserNotFound.Error(), err)
	case errors.Is(err, models.ErrUserAlreadyEnrolled):
		Error(w, r, http.StatusConflict, models.ErrUserAlreadyEnrolled.Error(), err)
	case errors.Is(err, models.ErrAuthorNotFound):
		Error(w, r, http.StatusBadRequest, models.ErrAuthorNotFound.Error(), err)
	case errors.Is(err, models.ErrMentorNotFound):
		Error(w, r, http.StatusBadRequest, models.ErrMentorNotFound.Error(), err)
	default:
		Error(w, r, http.StatusInternalServerError, "internal server error", err)
	}
}
