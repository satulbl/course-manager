package mentor

import (
	"context"
	"course-manager/internal/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

const columns = `u.id, u.email, u.password_hash, u.name, u.role_id, r.code AS role_code, r.name as role_name, u.created_at, u.updated_at`

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return r.get(ctx,
		`SELECT `+columns+`
		 FROM users u
		 JOIN roles r ON r.id = u.role_id
		 WHERE u.id = $1 AND u.role_id = $2`,
		id, models.RoleTeacher,
	)
}

func (r *Repository) get(ctx context.Context, query string, arg ...any) (*models.User, error) {
	rows, err := r.db.Query(ctx, query, arg...)
	if err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}

	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan mentor: %w", err)
	}
	user := m.toModel()
	return &user, nil
}
