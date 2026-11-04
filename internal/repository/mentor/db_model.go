package mentor

import (
	"course-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

type dbModel struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Name         string    `db:"name"`
	RoleID       int64     `db:"role_id"`
	RoleCode     string    `db:"role_code"`
	RoleName     string    `db:"role_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (m dbModel) toModel() models.User {
	return models.User{
		ID:           m.ID,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		Name:         m.Name,
		CreatedAt:    m.CreatedAt,
		Role: models.UserRole{
			Id:   m.RoleID,
			Name: m.RoleName,
			Code: m.RoleCode,
		},
	}
}