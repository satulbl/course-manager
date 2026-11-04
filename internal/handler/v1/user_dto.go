package v1

import (
	"time"

	"course-manager/internal/models"

	"github.com/google/uuid"
)

// userResponse — пользователь в ответе API. Пароль и role_id наружу не отдаём.
type userRole struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type userResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      userRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(u *models.User) userResponse {
	return userResponse{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
		Role: userRole{
			Name: u.Role.Name,
			Code: u.Role.Code,
		},
		CreatedAt: u.CreatedAt,
	}
}

type usersListResponse struct {
	Users      []userResponse `json:"users"`
	Limit      int            `json:"limit"`
	Offset     int            `json:"offset"`
	TotalCount int64          `json:"totalCount"`
}

func newUserListResponse(users []models.User, limit, offset int, total int64) usersListResponse {
	items := make([]userResponse, 0, len(users))
	for _, u := range users {
		items = append(items, newUserResponse(&u))
	}

	return usersListResponse{
		Users:      items,
		Limit:      limit,
		Offset:     offset,
		TotalCount: total,
	}
}
