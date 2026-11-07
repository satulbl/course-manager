package course

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
	fkViolation = "23503"
	fkCoursesMentor = "courses_mentor_id_fkey"
)
const selectColumns = `id, title, description, price, author_id, mentor_id, start_date, end_date, created_at, updated_at`

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func isMentorFKViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) &&
		pgErr.Code == fkViolation &&
		pgErr.ConstraintName == fkCoursesMentor
}

func (r *Repository) Create(ctx context.Context, c *models.Course) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO courses (id, title, description, price, author_id, mentor_id, start_date, end_date)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		c.ID, c.Title, c.Description, c.Price, c.AuthorID, c.MentorID, c.StartDate, c.EndDate)
	if err != nil {
		if isMentorFKViolation(err) {
			return models.ErrMentorNotFound
		}
		return fmt.Errorf("insert course: %w", err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, c *models.Course) error {
	tag, err := r.db.Exec(ctx,
		`UPDATE courses
		 SET title = $2, description = $3, price = $4, mentor_id = $5, start_date = $6, end_date = $7, updated_at = $8
		 WHERE id = $1`,
		c.ID, c.Title, c.Description, c.Price, c.MentorID, c.StartDate, c.EndDate, c.UpdatedAt,
	)
	if err != nil {
		if isMentorFKViolation(err) {
			return models.ErrMentorNotFound
		}
		return fmt.Errorf("update course: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	tag, err := r.db.Exec(ctx, `DELETE FROM courses WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete course: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return models.ErrNotFound
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*models.Course, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+selectColumns+`
	     FROM courses
	     WHERE id = $1`,
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("select course by id: %w", err)
	}

	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan course: %w", err)
	}
	course := m.toModel()
	return &course, nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.Course, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+selectColumns+`
         FROM courses
         ORDER BY created_at DESC
         LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select courses: %w", err)
	}

	return collectCourses(rows)
}

func (r *Repository) Count(ctx context.Context) (int64, error) {
	var total int64
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM courses`).Scan(&total); err != nil {
		return 0, fmt.Errorf("count courses: %w", err)
	}
	return total, nil
}

func collectCourses(rows pgx.Rows) ([]models.Course, error) {
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		return nil, fmt.Errorf("scan courses: %w", err)
	}

	courses := make([]models.Course, 0, len(items))
	for _, c := range items {
		courses = append(courses, c.toModel())
	}
	return courses, nil
}
