package usercourse

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

	fkUserCoursesUser   = "user_courses_user_id_fkey"
	fkUserCoursesCourse = "user_courses_course_id_fkey"
	uqUserCourses       = "user_courses_course_id_key"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserCourses(ctx context.Context, f models.UserCourseFilter) ([]models.UserCourse, error) {
	rows, err := r.db.Query(ctx,
		`SELECT uc.id,
		    uc.is_approved,
		    u.id    AS user_id,
			u.email AS user_email,
			u.name  AS user_name,
			c.id    AS course_id,
			c.title AS course_title
	FROM user_courses uc
	JOIN users u ON u.id = uc.user_id
	JOIN courses c ON c.id = uc.course_id
	WHERE ($1::uuid IS NULL OR uc.user_id = $1)
	 AND ($2::boolean IS NULL OR uc.is_approved = $2)
	ORDER BY uc.created_at DESC`,
		f.UserID, f.IsApproved,
	)
	if err != nil {
		return nil, fmt.Errorf("select user courses: %w", err)
	}

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		return nil, fmt.Errorf("scan user courses: %w", err)
	}

	userCourses := make([]models.UserCourse, 0, len(items))
	for _, uc := range items {
		userCourses = append(userCourses, uc.toModel())
	}

	if len(userCourses) == 0 && f.UserID != nil {
		var exists bool
		err := r.db.QueryRow(ctx,
			`SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)`,
			*f.UserID,
		).Scan(&exists)
		if err != nil {
			return nil, fmt.Errorf("check user exists: %w", err)
		}
		if !exists {
			return nil, models.ErrUserNotFound
		}
	}

	return userCourses, nil
}

func (r *Repository) Create(ctx context.Context, uc *models.UserCourse) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO user_courses 
	     (id, user_id, course_id)
	     VALUES ($1, $2, $3)`,
		uc.ID, uc.User.ID, uc.Course.ID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case pgErr.Code == fkViolation && pgErr.ConstraintName == fkUserCoursesUser:
				return models.ErrUserNotFound
			case pgErr.Code == fkViolation && pgErr.ConstraintName == fkUserCoursesCourse:
				return models.ErrCourseNotFound
			case pgErr.Code == uniqueViolation && pgErr.ConstraintName == uqUserCourses:
				return models.ErrUserAlreadyEnrolled
			}
		}
		return fmt.Errorf("insert user course: %w", err)
	}

	return nil
}

func (r *Repository) Approve(ctx context.Context, id uuid.UUID) error {
	res, err := r.db.Exec(ctx,
		`UPDATE user_courses
	 SET is_approved = true, updated_at = now()
	 WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("approve user course: %w", err)
	}
	if res.RowsAffected() == 0 {
		return models.ErrNotFound
	}

	return nil
}
