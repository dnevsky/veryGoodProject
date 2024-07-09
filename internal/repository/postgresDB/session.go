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

type SessionRepository struct {
	db *pgxpool.Pool
}

func NewSessionRepository(db *pgxpool.Pool) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Create(ctx context.Context, uid uint64, ip string) (session *models.Session, err error) {
	session = &models.Session{}
	query := `INSERT INTO sessions (uid, ip) VALUES ($1, $2) RETURNING id, uid, ip, created_at`

	err = r.db.QueryRow(ctx, query, uid, ip).Scan(&session.Id, &session.Uid, &session.Ip, &session.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return session, nil
}

func (r *SessionRepository) FindById(ctx context.Context, id string) (session *models.Session, err error) {
	session = &models.Session{}
	query := `SELECT id, uid, created_at FROM sessions WHERE id=$1 ORDER BY created_at DESC`

	row := r.db.QueryRow(ctx, query, id)
	err = row.Scan(&session.Id, &session.Uid, &session.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to scan session: %w", err)
	}

	return session, nil
}

func (r *SessionRepository) DeleteTokens(ctx context.Context, uid uint64) (err error) {
	query := `DELETE FROM sessions WHERE uid=$1`
	_, err = r.db.Exec(ctx, query, uid)
	return err
}
