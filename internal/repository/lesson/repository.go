package lesson

import (
	"context"
	"course-manager/internal/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	fkViolation     = "23503"
	uniqueViolation = "23505"
)
const selectLesson = `SELECT 
    l.id, 
    l.title, 
    l.description, 
    l.content, 
    l.lesson_number, 
    l.course_id, 
    c.title AS course_title, 
    c.mentor_id AS mentor_id, 
    u.name AS mentor_name, 
    l.created_at, 
    l.updated_at 
FROM lessons l 
JOIN courses c ON l.course_id = c.id
JOIN users u ON c.mentor_id = u.id`

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, l *models.Lesson) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO lessons (id, title, description, content, lesson_number, course_id)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		l.ID, l.Title, l.Description, l.Content, l.LessonNumber, l.Course.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case pgErr.Code == fkViolation && pgErr.ConstraintName == "lessons_course_id_fkey":
				return models.ErrCourseNotFound
			case pgErr.Code == uniqueViolation && pgErr.ConstraintName == "lessons_course_id_lesson_number_key":
				return models.ErrLessonNumberTaken
			}
		}
		return fmt.Errorf("insert lesson: %w", err)
	}
	return nil
}

func (r *Repository) GetById(ctx context.Context, id uuid.UUID) (*models.Lesson, error) {
	rows, err := r.db.Query(ctx, selectLesson+` WHERE l.id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("select lesson by id: %w", err)
	}

	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan lesson: %w", err)
	}

	lesson := m.toModel()
	return &lesson, nil
}

func (r *Repository) Count(ctx context.Context) (int64, error) {
	var total int64
	err := r.db.QueryRow(ctx, "SELECT COUNT(*) FROM lessons").Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("count lessons: %w", err)
	}

	return total, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.Lesson, error) {
	rows, err := r.db.Query(ctx, selectLesson+` ORDER BY l.created_at DESC LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("select lessons: %w", err)
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		return nil, fmt.Errorf("scan lessons: %w", err)
	}

	lessons := make([]models.Lesson, 0, len(items))
	for _, l := range items {
		lessons = append(lessons, l.toModel())
	}

	return lessons, nil
}
