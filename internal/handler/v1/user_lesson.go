package v1

import (
	"course-manager/internal/handler/httpx"
	"course-manager/internal/handler/middleware"
	"net/http"
)

func (h *Handler) StartUserLesson(w http.ResponseWriter, r *http.Request) {
	var req *userLessonRequest
	if ok := httpx.DecodeJSON(w, r, &req); !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	id, err := h.services.UserLesson.Create(r.Context(), user.ID, req.toInput())
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newSimpleIdResponse(id))
}

func (h *Handler) CompleteUserLesson(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.ParseID(w, r)
	if !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	err := h.services.UserLesson.Complete(r.Context(), user, id)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
