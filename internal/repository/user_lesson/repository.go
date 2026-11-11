package userlessons

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

	fkUserLessonsUser = "user_lessons_user_id_fkey"
	fkUserLessons     = "user_lessons_lesson_id_fkey"
	uqUserLesson      = "user_lessons_user_id_lesson_id_key"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetById(ctx context.Context, id uuid.UUID) (*models.UserLesson, error) {
	rows, err := r.db.Query(ctx,
		`SELECT ul.id,
		        ul.completed_at,
		        ul.created_at,
		        u.id           AS user_id,
		        u.name         AS user_name,
		        u.email        AS user_email,
		        l.id           AS lesson_id,
		        l.title        AS lesson_title,
		        l.lesson_number,
		        c.id           AS course_id,
		        c.title        AS course_title
		 FROM user_lessons ul
		 JOIN users u   ON u.id = ul.user_id
		 JOIN lessons l ON l.id = ul.lesson_id
		 JOIN courses c ON c.id = l.course_id
		 WHERE ul.id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("select user lesson: %w", err)
	}

	item, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNotFound
		}
		return nil, fmt.Errorf("scan user lesson: %w", err)
	}

	ul := item.toModel()
	return &ul, nil
}

func (r *Repository) Complete(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.Exec(ctx,
		`UPDATE user_lessons 
	SET completed_at = now()
	WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("update user lesson: %w", err)
	}
	if res.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (r *Repository) Create(ctx context.Context, ul *models.UserLesson) error {
	res, err := r.db.Exec(ctx,
		`INSERT INTO user_lessons (id, user_id, lesson_id)
		 SELECT $1::uuid, $2::uuid, l.id
		 FROM lessons l
		 JOIN user_courses uc ON uc.course_id = l.course_id
		                     AND uc.user_id = $2::uuid
		                     AND uc.is_approved
		 WHERE l.id = $3::uuid`,
		ul.ID, ul.User.ID, ul.Lesson.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgErr.Code == uniqueViolation &&
			pgErr.ConstraintName == uqUserLesson {
			return models.ErrUserAlreadyStartedLesson
		}
		return fmt.Errorf("insert user lesson: %w", err)
	}

	if res.RowsAffected() == 0 {
		var lessonExists bool
		if err := r.db.QueryRow(ctx,
			`SELECT EXISTS (SELECT 1 FROM lessons WHERE id = $1)`,
			ul.Lesson.ID,
		).Scan(&lessonExists); err != nil {
			return fmt.Errorf("check lesson exists: %w", err)
		}
		if !lessonExists {
			return models.ErrLessonNotFound
		}
		return models.ErrCourseNotApproved
	}

	return nil
}