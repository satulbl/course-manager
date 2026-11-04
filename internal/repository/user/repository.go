package user

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

const uniqueViolation = "23505"

const selectUser = `
	SELECT u.id, u.email, u.password_hash, u.name, u.role_id, r.code AS role_code, r.name as role_name, u.created_at, updated_at
	  FROM users u
	  JOIN roles r ON r.id = u.role_id`

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, u *models.User) error {
	err := r.db.QueryRow(ctx,
		`INSERT INTO users (id, email, password_hash, name)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at`,
		u.ID, u.Email, u.PasswordHash, u.Name,
	).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
			return models.ErrEmailTaken
		}
		return fmt.Errorf("insert user: %w", err)
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return r.get(ctx, selectUser+` WHERE u.id = $1`, id)
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return r.get(ctx, selectUser+` WHERE u.email = $1`, email)
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]models.User, error) {
	rows, err := r.db.Query(ctx,
		selectUser+` ORDER BY u.created_at DESC, u.id
		LIMIT $1 OFFSET $2`,
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}

	return collectUsers(rows)
}

func (r *Repository) Count(ctx context.Context) (int64, error) {
	var total int64
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM users`).Scan(&total); err != nil {
		return 0, fmt.Errorf("count users: %w", err)
	}
	return total, nil
}

func (r *Repository) get(ctx context.Context, query string, arg any) (*models.User, error) {
	rows, err := r.db.Query(ctx, query, arg)
	if err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}

	m, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbModel])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan user: %w", err)
	}
	user := m.toModel()
	return &user, nil
}

func collectUsers(rows pgx.Rows) ([]models.User, error) {
	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbModel])
	if err != nil {
		return nil, fmt.Errorf("scan users: %w", err)
	}

	users := make([]models.User, 0, len(items))
	for _, u := range items {
		users = append(users, u.toModel())
	}
	return users, nil
}
