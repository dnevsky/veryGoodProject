package v1

import (
	authDto "github.com/dnevsky/veryGoodProject/internal/dto/auth"
	"github.com/dnevsky/veryGoodProject/transport/rest/response"
	"github.com/gin-gonic/gin"
	"time"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (h *Handler) initAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", h.login)
	}
}

// @Summary Авторизация пользователя
// @Tags API для авторизации пользователей
// @Description Авторизация пользователя
// @ID auth-login
// @Accept json
// @Produce json
// @Param input body auth.AuthDTO true "credentials"
// @Success 200 {object} tokenResponse "Токен"
// @Failure default {object} response.Data
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input authDto.AuthDTO

	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	if err := input.Validate(); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	input.Ip = c.Request.RemoteAddr

	resp, err := h.services.User.Login(c, input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	accessTokenTTL, err := time.ParseDuration(h.config.Auth.AccessTokenTTL)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: tokenResponse{
			AccessToken: resp.AccessToken,
			ExpiresIn:   int64(accessTokenTTL.Minutes()),
		},
	})
}
