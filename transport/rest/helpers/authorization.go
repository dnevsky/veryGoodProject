package helpers

import (
	"fmt"
	"github.com/dnevsky/veryGoodProject/internal/models"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	UserCtx             = "userId"
)

func (m *Manager) GetUserIdAuthorization(c *gin.Context) (uint64, error) {
	if userId, exists := c.Get(UserCtx); exists {
		userId, ok := userId.(uint64)
		if !ok {
			return 0, fmt.Errorf("invalid user id type")
		}

		return userId, nil
	}
	return 0, models.ErrUnauthorized
}
