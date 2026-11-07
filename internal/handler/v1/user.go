package v1

import (
	httpx "course-manager/internal/handler/httpx"
	"net/http"
)

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	limit, ok := httpx.QueryInt(w, r, "limit", defaultListLimit)
	if !ok {
		return
	}

	offset, ok := httpx.QueryInt(w, r, "offset", defaultListOffset)
	if !ok {
		return
	}

	users, total, err := h.services.User.List(r.Context(), limit, offset)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newUserListResponse(users, limit, offset, total))
}
