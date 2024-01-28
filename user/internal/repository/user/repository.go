package user

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
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

	emailExists, err := u.emailExists(ctx, user.Email)
	if err != nil {
		return err
	}

	if emailExists {
		return userErrors.ErrEmailAlreadyExists
	}

	sqlQuery := `
		INSERT INTO users (id, name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = u.db.Exec(ctx, sqlQuery, user.ID, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		u.log.Error("error creating user", zap.Error(err), zap.String("op", op))
		return err
	}

	return nil
}

func (u *userRepo) emailExists(ctx context.Context, email string) (bool, error) {
	const op = "userRepo.emailExists"

	sqlQuery := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	err := u.db.QueryRow(ctx, sqlQuery, email).Scan(&exists)
	if err != nil {
		u.log.Error("error checking if email exists", zap.Error(err), zap.String("op", op))
		return false, err
	}

	return exists, nil
}

func (u *userRepo) Get(ctx context.Context, id string) (*model.User, error) {
	const op = "userRepo.Get"

	sqlQuery := `SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1`

	rows, err := u.db.Query(ctx, sqlQuery, id)
	if err != nil {
		u.log.Error("error getting user", zap.Error(err), zap.String("op", op))
		return nil, err
	}

	var user repoModel.User

	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		u.log.Error("error scanning user", zap.Error(err), zap.String("op", op))
		return nil, err
	}

	return dao.ToUserFromRepo(&user), nil
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	const op = "userRepo.Update"

	sqlQuery := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, updated_at = $5 WHERE id = $6`

	result, err := u.db.Exec(ctx, sqlQuery, user.Name, user.Email, user.Password, user.Role, user.UpdatedAt, user.ID)
	if err != nil {
		u.log.Error("error updating user", zap.Error(err), zap.String("op", op))
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows were affected, user with ID %s not found", user.ID)
		u.log.Error("error updating user", zap.Error(err), zap.String("op", op))
		return err
	}

	return nil
}

func (u *userRepo) Delete(ctx context.Context, id string) error {
	const op = "userRepo.Delete"

	sqlQuery := `DELETE FROM users WHERE id = $1`

	result, err := u.db.Exec(ctx, sqlQuery, id)
	if err != nil {
		u.log.Error("error deleting user", zap.Error(err), zap.String("op", op))
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		err := fmt.Errorf("no rows were affected, user with ID %s not found", id)
		u.log.Error("error deleting user", zap.Error(err), zap.String("op", op))
		return err
	}

	return nil
}
