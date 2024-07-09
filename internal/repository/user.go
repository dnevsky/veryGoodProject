package repository

import (
	"context"
	"github.com/dnevsky/veryGoodProject/internal/models"
)

type UserRepository interface {
	FindByLoginAndPass(ctx context.Context, login, passwordMd5 string) (user *models.User, err error)
}
