package middleware

import (
	"errors"
	"github.com/dnevsky/veryGoodProject/internal/models"
	repoErrors "github.com/dnevsky/veryGoodProject/internal/repository/errors"
	"github.com/dnevsky/veryGoodProject/internal/service"
	"github.com/dnevsky/veryGoodProject/transport/rest/helpers"
	"github.com/dnevsky/veryGoodProject/transport/rest/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthUser(sessionService service.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := parseAuthHeader(c)
		if err != nil {
			response.NewResponse(c, http.StatusUnauthorized, "unauthorized")
			return
		}

		session, err := sessionService.FindById(c, token)
		if err != nil && errors.Is(err, repoErrors.ErrNotFound) {
			response.NewResponse(c, http.StatusUnauthorized, "unauthorized")
		}

		err = sessionService.VerifySession(session)
		if err != nil {
			response.NewResponse(c, http.StatusUnauthorized, "token expired")
		}

		c.Set(helpers.UserCtx, session.Uid)
		c.Next()
	}
}

func parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(helpers.AuthorizationHeader)
	if header == "" {
		return "", models.ErrEmptyAuthHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", models.ErrInvalidAuthHeader
	}
	if len(headerParts[1]) == 0 {
		return "", models.ErrTokenIsEmpty
	}

	return headerParts[1], nil
}
