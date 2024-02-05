package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"strings"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/repository/user/dao"
	repoModel "user/internal/repository/user/model"
	userErrors "user/pkg/user_errors"
)

type userRepo struct {
	db *pgxpool.Pool

	log *zap.Logger
}

func NewUserRepository(db *pgxpool.Pool, log *zap.Logger) repository.UserRepository {
	return &userRepo{
		db:  db,
		log: log,
	}
}

func (u *userRepo) Create(ctx context.Context, user *model.User) error {
	const op = "userRepo.Create"
	sqlQuery := `
		INSERT INTO users (id, name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := u.db.Exec(ctx, sqlQuery, user.ID, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return handleEmailConflictError(err)
	}

	return nil
}

func handleEmailConflictError(err error) error {
	if err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return userErrors.ErrEmailAlreadyExists
	}
	return err
}

func (u *userRepo) GetById(ctx context.Context, id string) (*model.User, error) {
	sqlQuery := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1`

	var user repoModel.User

	err := u.db.QueryRow(ctx, sqlQuery, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, userErrors.ErrWrongEmailOrPassword
		}

		return nil, err
	}

	return dao.ToUserFromRepo(&user), nil
}

func (u *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	sqlQuery := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1`

	var user repoModel.User

	err := u.db.QueryRow(ctx, sqlQuery, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, userErrors.ErrWrongEmailOrPassword
		}

		return nil, err
	}

	return dao.ToUserFromRepo(&user), nil
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	sqlQuery := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, updated_at = $5 WHERE id = $6`

	_, err := u.db.Exec(ctx, sqlQuery, user.Name, user.Email, user.Password, user.Role, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) UpdateRefreshToken(ctx context.Context, id string, token string) error {
	sqlQuery := `UPDATE users SET refresh_token = $1 WHERE id = $2`

	_, err := u.db.Exec(ctx, sqlQuery, token, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	sqlQuery := `DELETE FROM users WHERE id = $1`

	result, err := u.db.Exec(ctx, sqlQuery, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows were affected, user with ID %s not found", id)
		return err
	}

	return nil
}
