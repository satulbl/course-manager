package v1

import (
	"course-manager/internal/handler/httpx"
	"course-manager/internal/handler/middleware"
	"course-manager/internal/models"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (h *Handler) UserCourses(w http.ResponseWriter, r *http.Request) {
	filter := models.UserCourseFilter{}

	if rawID := r.URL.Query().Get("userId"); rawID != "" {
		id, err := uuid.Parse(rawID)
		if err != nil {
			httpx.Error(w, r, http.StatusBadRequest, "invalid userId", err)
			return
		}
		filter.UserID = &id
	}

	if approved := r.URL.Query().Get("approved"); approved != "" {
		isApproved, err := strconv.ParseBool(approved)
		if err != nil {
			httpx.Error(w, r, http.StatusBadRequest, "approve must be true or false", err)
			return
		}
		filter.IsApproved = &isApproved
	}
	user := middleware.CurrentUser(r.Context())

	uc, err := h.services.UserCourse.GetUserCourses(r.Context(), user, filter)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newUserCoursesResponse(uc))
}

func (h *Handler) CreateUserCourse(w http.ResponseWriter, r *http.Request) {
	var req *createUserCourseRequest
	ok := httpx.DecodeJSON(w, r, &req)
	if !ok {
		return
	}

	id, err := h.services.UserCourse.Create(r.Context(), req.toInput())
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newSimpleIdResponse(id))
}

func (h *Handler) ApproveUserCourse(w http.ResponseWriter, r *http.Request) {
	id, ok := httpx.ParseID(w, r)
	if !ok {
		return
	}

	user := middleware.CurrentUser(r.Context())
	err := h.services.UserCourse.Approve(r.Context(), user, id)
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
