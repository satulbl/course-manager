package usercourse

import (
	"course-manager/internal/models"

	"github.com/google/uuid"
)

type Course struct {
	ID    uuid.UUID `db:"course_id"`
	Title string    `db:"course_title"`
}

type User struct {
	ID    uuid.UUID `db:"user_id"`
	Name  string    `db:"user_name"`
	Email string    `db:"user_email"`
}

type dbModel struct {
	ID uuid.UUID `db:"id"`
	User
	Course
	IsApproved bool `db:"is_approved"`
}

func (m dbModel) toModel() models.UserCourse {
	return models.UserCourse{
		ID: m.ID,
		User: models.ApplicationUser{
			ID:    m.User.ID,
			Name:  m.User.Name,
			Email: m.User.Email,
		},
		Course: models.ApplicationCourse{
			ID:    m.Course.ID,
			Title: m.Course.Title,
		},
		IsApproved: m.IsApproved,
	}
}
