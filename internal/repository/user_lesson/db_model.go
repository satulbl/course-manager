package userlessons

import (
	"course-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `db:"user_id"`
	Name  string    `db:"user_name"`
	Email string    `db:"user_email"`
}

type Course struct {
	ID    uuid.UUID `db:"course_id"`
	Title string    `db:"course_title"`
}

type Lesson struct {
	ID           uuid.UUID `db:"lesson_id"`
	Title        string    `db:"lesson_title"`
	LessonNumber int       `db:"lesson_number"`
	Course
}

type dbModel struct {
	ID          uuid.UUID  `db:"id"`
	CompletedAt *time.Time `db:"completed_at"`
	CreatedAt   time.Time  `db:"created_at"`
	User
	Lesson
}

func (m dbModel) toModel() models.UserLesson {
	return models.UserLesson{
		ID: m.ID,
		User: models.ApplicationUser{
			ID:    m.User.ID,
			Name:  m.User.Name,
			Email: m.User.Email,
		},
		Lesson: models.ApplicationLesson{
			ID:           m.Lesson.ID,
			Title:        m.Lesson.Title,
			LessonNumber: m.Lesson.LessonNumber,
			Course: models.ApplicationCourse{
				ID:    m.Lesson.Course.ID,
				Title: m.Lesson.Course.Title,
			},
		},
		CompletedAt: m.CompletedAt,
		CreatedAt:   m.CreatedAt,
	}
}
