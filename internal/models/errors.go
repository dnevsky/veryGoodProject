package models

import "errors"

var (
	ErrUnauthorized          = errors.New("unauthorized")
	ErrInvalidRequestParams  = errors.New("invalid request parameters. see invalidArgs")
	ErrBrokenPipe            = errors.New("write: broken pipe")
	ErrConnectionResetByPeer = errors.New("connection reset by peer")
	ErrBadLoginOrPassword    = errors.New("bad login or password")
	ErrEmptyAuthHeader       = errors.New("empty auth header")
	ErrInvalidAuthHeader     = errors.New("invalid auth header")
	ErrTokenIsEmpty          = errors.New("token is empty")
	ErrTokenExpired          = errors.New("token expired")
	ErrBadAssetName          = errors.New("bad asset name")
	ErrEmptyBody             = errors.New("empty body")
)

type ErrType int32

const (
	ErrTypeDefault    ErrType = 0
	ErrTypeValidation ErrType = 1
)
