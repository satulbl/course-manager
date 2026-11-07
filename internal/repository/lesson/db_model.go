package lesson

import (
	"course-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID    uuid.UUID `db:"course_id"`
	Title string    `db:"course_title"`
}

type Mentor struct {
	ID   uuid.UUID `db:"mentor_id"`
	Name string    `db:"mentor_name"`
}

type dbModel struct {
	ID           uuid.UUID `db:"id"`
	Title        string    `db:"title"`
	Description  string    `db:"description"`
	Content      string    `db:"content"`
	LessonNumber int       `db:"lesson_number"`
	Course
	Mentor
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (m dbModel) toModel() models.Lesson {
	return models.Lesson{
		ID:           m.ID,
		Title:        m.Title,
		Description:  m.Description,
		Content:      m.Content,
		LessonNumber: m.LessonNumber,
		Course: models.LessonCourse{
			ID:    m.Course.ID,
			Title: m.Course.Title,
		},
		Mentor: models.LessonMentor{
			ID:   m.Mentor.ID,
			Name: m.Mentor.Name,
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
