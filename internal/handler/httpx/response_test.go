package httpx

import (
	"course-manager/internal/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceErrorMapsUserLessonErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{name: "lesson not found", err: fmt.Errorf("create user lesson: %w", models.ErrLessonNotFound), want: http.StatusBadRequest},
		{name: "lesson already started", err: fmt.Errorf("create user lesson: %w", models.ErrUserAlreadyStartedLesson), want: http.StatusConflict},
		{name: "lesson already completed", err: fmt.Errorf("complete user lesson: %w", models.ErrAlreadyCompleted), want: http.StatusConflict},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/lesson/start", nil)
			rec := httptest.NewRecorder()

			ServiceError(rec, req, tt.err)

			if rec.Code != tt.want {
				t.Fatalf("ServiceError returned status %d, want %d", rec.Code, tt.want)
			}
		})
	}
}
