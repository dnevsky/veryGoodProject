package repository

import (
	"context"
	"github.com/dnevsky/veryGoodProject/internal/models"
)

type SessionRepository interface {
	Create(ctx context.Context, uid uint64, ip string) (session *models.Session, err error)
	FindById(ctx context.Context, id string) (session *models.Session, err error)
	DeleteTokens(ctx context.Context, uid uint64) (err error)
}
