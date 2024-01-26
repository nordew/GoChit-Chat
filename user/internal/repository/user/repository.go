package user

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/repository/user/dao"
	repoModel "user/internal/repository/user/model"
	queryBuilder "user/pkg/query_builder"
)

type userRepo struct {
	db *pgxpool.Pool

	filter queryBuilder.QueryBuilder
}

func NewUserRepository(db *pgxpool.Pool, filter queryBuilder.QueryBuilder) repository.UserRepository {
	return &userRepo{
		db:     db,
		filter: filter,
	}
}

func (u *userRepo) Create(ctx context.Context, user *model.User) error {
	sqlQuery := `INSERT INTO users (id, name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	rows, err := u.db.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}

	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Get(ctx context.Context, filter *repository.GetFilter) (*model.User, error) {
	sqlQuery, _ := u.filter.SetBaseSQL(`SELECT id, name, email, password, refresh_token, role, created_at, updated_at FROM users`).
		SetWhere(`ID`, filter.ID).
		SetWhere(`Email`, filter.Email).
		Generate()

	var user repoModel.User

	rows, err := u.db.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		return nil, err
	}

	return dao.ToUserFromRepo(&user), nil
}

func (u *userRepo) Update(ctx context.Context, user *model.User) error {
	sqlQuery := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, updated_at = $5 WHERE id = $6`

	rows, err := u.db.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}

	err = pgxscan.ScanOne(&user, rows)
	if err != nil {
		return err
	}

	return nil
}
