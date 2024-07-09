package repository

import (
	"github.com/dnevsky/veryGoodProject/internal/repository/postgresDB"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repositories struct {
	UserRepository    UserRepository
	SessionRepository SessionRepository
	AssetRepository   AssetRepository
	// ...
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		UserRepository:    postgresDB.NewUserRepository(db),
		SessionRepository: postgresDB.NewSessionRepository(db),
		AssetRepository:   postgresDB.NewAssetRepository(db),
		// ...
	}
}
