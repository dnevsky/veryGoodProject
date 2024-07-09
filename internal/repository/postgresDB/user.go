package postgresDB

import (
	"context"
	"errors"
	"fmt"
	"github.com/dnevsky/veryGoodProject/internal/models"
	repoErrors "github.com/dnevsky/veryGoodProject/internal/repository/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByLoginAndPass(ctx context.Context, login, passwordMd5 string) (user *models.User, err error) {
	user = &models.User{}
	query := `SELECT id, login, password_hash, created_at FROM users WHERE login=$1 AND password_hash=$2`

	row := r.db.QueryRow(ctx, query, login, passwordMd5)
	err = row.Scan(&user.Id, &user.Login, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	return user, nil
}
