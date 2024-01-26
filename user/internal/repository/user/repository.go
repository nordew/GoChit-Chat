package user

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"user/internal/model"
	"user/internal/repository"
	"user/internal/repository/user/dao"
	repoModel "user/internal/repository/user/model"
	queryBuilder "user/pkg/query_builder"
)

type userRepo struct {
	db *pgxpool.Conn

	filter queryBuilder.QueryBuilder
}

func NewUserRepository(db *pgxpool.Conn, filter queryBuilder.QueryBuilder) repository.UserRepository {
	return &userRepo{
		db:     db,
		filter: filter,
	}
}

func (u *userRepo) Create(ctx context.Context, user *model.User) error {
	sqlQuery := `INSERT INTO users (id, name, email, password, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := u.db.Exec(ctx, sqlQuery, user.ID, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Get(ctx context.Context, filter *repository.GetFilter) (*model.User, error) {
	sqlQuery, args, err := u.filter.BuildQuery(filter)
	if err != nil {
		return nil, err
	}

	var user repoModel.User
	err = u.db.QueryRow(ctx, sqlQuery, args...).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RefreshToken, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
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
