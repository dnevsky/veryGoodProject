package service

import (
	"github.com/dnevsky/veryGoodProject/internal/configs"
	"github.com/dnevsky/veryGoodProject/internal/repository"
	"github.com/dnevsky/veryGoodProject/pkg/logger"
)

type Services struct {
	Repositories *repository.Repositories
	Logger       logger.Logger
	User         User
	Session      Session
	Asset        Asset
	// ...
}

type Deps struct {
	Repositories *repository.Repositories
	Logger       logger.Logger
	Config       configs.Config
}

func NewServices(deps Deps) (*Services, error) {
	userService := NewUserService(
		deps.Repositories.UserRepository,
		deps.Repositories.SessionRepository,
		UserServiceConfig{
			AccessTokenTTL: deps.Config.Auth.AccessTokenTTL,
		},
	)

	sessionService := NewSessionService(
		deps.Repositories.SessionRepository,
		SessionServiceConfig{
			AccessTokenTTL: deps.Config.Auth.AccessTokenTTL,
		},
	)

	assetService := NewAssetService(
		deps.Repositories.AssetRepository,
	)

	return &Services{
		Repositories: deps.Repositories,
		Logger:       deps.Logger,
		User:         userService,
		Session:      sessionService,
		Asset:        assetService,
	}, nil
}
