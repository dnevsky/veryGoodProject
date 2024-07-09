package helpers

import (
	"github.com/dnevsky/veryGoodProject/internal/repository"
	"github.com/dnevsky/veryGoodProject/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Helpers interface {
	GetUserIdAuthorization(c *gin.Context) (uint64, error)
	BindData(c *gin.Context, req interface{}) error
	ErrorsHandle(c *gin.Context, err error)
	LogError(err error)
	GetIdFromPath(c *gin.Context, key string) (uint, error)
}

type Manager struct {
	UserRepository repository.UserRepository
	Logger         logger.Logger
}

func NewManager(
	logManager logger.Logger,
) *Manager {
	return &Manager{
		Logger: logManager,
	}
}
