package v1

import (
	"course-manager/internal/handler/httpx"
	"course-manager/internal/handler/middleware"
	"net/http"
)

func (h *Handler) CreateLesson(w http.ResponseWriter, r *http.Request) {
	var req lessonRequest
	if ok := httpx.DecodeJSON(w, r, &req); !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	id, err := h.services.Lesson.Create(r.Context(), user, req.toInput())
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newCreateLessonResponse(id))
}

func (h *Handler) GetLessonById(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.ParseID(w, r)
	if !ok {
		return
	}

	lesson, err := h.services.Lesson.GetById(r.Context(), id)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newLessonResponse(lesson))
}

func (h *Handler) ListLessons(w http.ResponseWriter, r *http.Request) {
	limit, ok := httpx.QueryInt(w, r, "limit", defaultListLimit)
	if !ok {
		return
	}

	offset, ok := httpx.QueryInt(w, r, "offset", defaultListOffset)
	if !ok {
		return
	}

	lessons, total, err := h.services.Lesson.List(r.Context(), limit, offset)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newLessonListResponse(lessons, total))
}
