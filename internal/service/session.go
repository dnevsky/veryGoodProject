package service

import (
	"context"
	"github.com/dnevsky/veryGoodProject/internal/models"
	"github.com/dnevsky/veryGoodProject/internal/repository"
	"time"
)

type SessionServiceConfig struct {
	AccessTokenTTL string
}

type Session interface {
	FindById(ctx context.Context, id string) (session *models.Session, err error)
	VerifySession(session *models.Session) (err error)
}

type SessionService struct {
	SessionRepository repository.SessionRepository

	config SessionServiceConfig
}

func NewSessionService(
	sessionRepo repository.SessionRepository,
	config SessionServiceConfig,
) *SessionService {
	return &SessionService{
		SessionRepository: sessionRepo,
		config:            config,
	}
}

func (s *SessionService) FindById(ctx context.Context, id string) (session *models.Session, err error) {
	return s.SessionRepository.FindById(ctx, id)
}

func (s *SessionService) VerifySession(session *models.Session) (err error) {
	tokenTTL, err := time.ParseDuration(s.config.AccessTokenTTL)
	if err != nil {
		return err
	}

	if time.Since(session.CreatedAt) > tokenTTL {
		return models.ErrTokenExpired
	}

	return nil
}
