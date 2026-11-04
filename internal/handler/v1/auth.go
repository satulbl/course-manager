package v1

import (
	"course-manager/internal/handler/middleware"
	"course-manager/internal/models"
	httpx "course-manager/internal/handler/httpx"

	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if !httpx.DecodeJSON(w, r, &req) {
		return
	}

	user, err := h.services.Auth.Register(r.Context(), req.toInput())
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusCreated, newUserResponse(user))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if !httpx.DecodeJSON(w, r, &req) {
		return
	}

	tokens, err := h.services.Auth.Login(r.Context(), models.LoginInput{
		Email:      req.Email,
		Password:   req.Password,
		DeviceType: models.DeviceType(r.Header.Get("X-Device-Type")),
		IP:         httpx.ClientIP(r),
		UserAgent:  r.UserAgent(),
	})
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newTokenPairResponse(tokens))
}

// Refresh принимает refresh-токен в заголовке Authorization: Bearer <refresh_token>.
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	tokens, err := h.services.Auth.Refresh(r.Context(), middleware.BearerToken(r))
	if err != nil {
		httpx.ServiceError(w, r, err)
		return
	}

	httpx.JSON(w, http.StatusOK, newTokenPairResponse(tokens))
}

// func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
// 	user := middleware.CurrentUser(r.Context())

// 	if err := h.services.Auth.Logout(r.Context(), user); err != nil {
// 		httpx.ServiceError(w, r, err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
