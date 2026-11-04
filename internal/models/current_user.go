package models

import "github.com/google/uuid"

// CurrentUser — пользователь, который выполняет запрос. Берётся из access-токена.
// Все правила прав — методы CurrentUser.
type CurrentUser struct {
	ID       uuid.UUID
	RoleID    int64
	SessionID string
}

func (u CurrentUser) CanCreate() bool {
	return writeRoles[u.RoleID]
}

func (u CurrentUser) CanEditCourse(authorID uuid.UUID) bool {
	if !u.CanCreate() {
		return false
	}
	return u.ID == authorID && u.RoleID == RoleTeacher || u.RoleID == RoleAdmin
}

