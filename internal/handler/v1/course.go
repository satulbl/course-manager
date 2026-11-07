package v1

import (
	"course-manager/internal/handler/httpx"
	"course-manager/internal/handler/middleware"
	"net/http"
)

const (
	defaultListLimit  = 20
	defaultListOffset = 0
)

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var req courseRequest
	if !httpx.DecodeJSON(w, r, &req) {
		return
	}

	user := middleware.CurrentUser(r.Context())
	err := h.services.Course.Create(r.Context(), user, req.toInput())
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusCreated, nil)
}

func (h *Handler) ListCourses(w http.ResponseWriter, r *http.Request) {
	limit, ok := httpx.QueryInt(w, r, "limit", defaultListLimit)
	if !ok {
		return
	}
	offset, ok := httpx.QueryInt(w, r, "offset", defaultListOffset)
	if !ok {
		return
	}

	courses, total, err := h.services.Course.List(r.Context(), limit, offset)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newCourseListResponse(courses, total))
}

func (h *Handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.ParseID(w, r)
	if !ok {
		return
	}

	var req courseRequest
	if !httpx.DecodeJSON(w, r, &req) {
		return
	}

	user := middleware.CurrentUser(r.Context())
	if err := h.services.Course.Update(r.Context(), user, id, req.toInput()); err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.ParseID(w, r)
	if !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	err := h.services.Course.Delete(r.Context(), user, id)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
