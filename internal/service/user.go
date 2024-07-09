package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	authDto "github.com/dnevsky/veryGoodProject/internal/dto/auth"
	"github.com/dnevsky/veryGoodProject/internal/models"
	"github.com/dnevsky/veryGoodProject/internal/repository"
	repoErrors "github.com/dnevsky/veryGoodProject/internal/repository/errors"
	"time"
)

type UserServiceConfig struct {
	AccessTokenTTL string
}

type User interface {
	Login(ctx context.Context, dto authDto.AuthDTO) (resp *authDto.SessionResponseDTO, err error)
}

type UserService struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository

	config UserServiceConfig
}

func NewUserService(
	userRepo repository.UserRepository,
	sessionRepo repository.SessionRepository,
	config UserServiceConfig,
) *UserService {
	return &UserService{
		UserRepository:    userRepo,
		SessionRepository: sessionRepo,
		config:            config,
	}
}

func (s *UserService) Login(ctx context.Context, dto authDto.AuthDTO) (resp *authDto.SessionResponseDTO, err error) {
	resp = &authDto.SessionResponseDTO{}
	passwordMd5 := md5.Sum([]byte(dto.Password))
	dto.Password = hex.EncodeToString(passwordMd5[:])

	user, err := s.UserRepository.FindByLoginAndPass(ctx, dto.Login, dto.Password)
	if err != nil && errors.Is(err, repoErrors.ErrNotFound) {
		return nil, models.ErrBadLoginOrPassword
	}
	if err != nil {
		return nil, err
	}

	err = s.SessionRepository.DeleteTokens(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	session, err := s.SessionRepository.Create(ctx, user.Id, dto.Ip)
	if err != nil {
		return nil, err
	}

	tokenTTL, err := time.ParseDuration(s.config.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	resp.AccessToken = session.Id
	resp.ExpiresIn = int64(tokenTTL.Minutes())

	return resp, nil
}
