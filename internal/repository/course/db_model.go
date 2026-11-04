package course

import (
	"course-manager/internal/models"
	"time"

	"github.com/google/uuid"
)

type dbModel struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Price       int       `db:"price"`
	Description string    `db:"description"`
	AuthorID    uuid.UUID `db:"author_id"`
	MentorId    uuid.UUID `db:"mentor_id"`
	StartDate   time.Time `db:"start_date"`
	EndDate     time.Time `db:"end_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (m dbModel) toModel() models.Course {
	return models.Course{
		ID:          m.ID,
		Title:       m.Title,
		Price:       m.Price,
		Description: m.Description,
		AuthorID:    m.AuthorID,
		MentorId:    m.MentorId,
		StartDate:   m.StartDate,
		EndDate:     m.EndDate,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
