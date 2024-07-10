package v1

import (
	assetDto "github.com/dnevsky/veryGoodProject/internal/dto/asset"
	"github.com/dnevsky/veryGoodProject/internal/models"
	"github.com/dnevsky/veryGoodProject/transport/rest/middleware"
	"github.com/dnevsky/veryGoodProject/transport/rest/response"
	"github.com/gin-gonic/gin"
	"io"
)

type uploadResp struct {
	Status string
}

func (h *Handler) initAssetRoutes(api *gin.RouterGroup) {
	asset := api.Group("/asset")
	{
		auth := asset.Group("/", middleware.AuthUser(h.services.Session))
		{
			auth.POST("/upload/:name", h.uploadAsset)
			auth.GET("/:name", h.getAsset)
			auth.GET("/", h.getAssets)
			auth.DELETE("/:name", h.deleteAsset)
		}

	}
}

// @Summary Upload файла в базу данных
// @Tags API для работы с ассетами
// @Description Аплоад файла на сервер
// @ID asset-upload
// @Accept json
// @Param Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Param name path string true "Название файла"
// @Param body body string true "Любое тело запроса"
// @Success 200 {object} uploadResp "Статус"
// @Failure default {object} response.Data
// @Router /asset/upload/{name} [post]
func (h *Handler) uploadAsset(c *gin.Context) {
	var input assetDto.UploadAssetDTO

	uid, err := h.helpers.GetUserIdAuthorization(c)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	name := c.Param("name")
	if len(name) < 3 || len(name) > 255 {
		h.helpers.ErrorsHandle(c, models.ErrBadAssetName)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}
	if len(body) == 0 {
		h.helpers.ErrorsHandle(c, models.ErrEmptyBody)
		return
	}

	input.XUserID = uid
	input.Name = name
	input.Body = body

	_, err = h.services.Asset.UploadAsset(c, input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: map[string]string{
			"status": "ok",
		},
	})
}

// @Summary Получение ассета
// @Tags API для работы с ассетами
// @Description Получение одного файла по его имени
// @ID asset-get
// @Accept json
// @Produce json
// @Param Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Param name path string true "Название файла"
// @Success 200 {object} asset.AssetResponseDTO "Ассет"
// @Failure default {object} response.Data
// @Router /asset/{name} [get]
func (h *Handler) getAsset(c *gin.Context) {
	var input assetDto.GetAssetDTO

	name := c.Param("name")
	if len(name) < 3 || len(name) > 255 {
		h.helpers.ErrorsHandle(c, models.ErrBadAssetName)
		return
	}

	uid, err := h.helpers.GetUserIdAuthorization(c)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	input.Name = name
	input.XUserID = uid

	asset, err := h.services.Asset.GetAsset(c, input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: asset.Data,
	})
}

// @Summary Получение списка ассетов
// @Tags API для работы с ассетами
// @Description Получение всех ассетов текущего пользователя
// @ID asset-getlist
// @Accept json
// @Produce json
// @Param Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Param input query asset.GetAssetsDTO true "пагинация"
// @Success 200 {object} []asset.AssetResponseDTO "Список ассетов"
// @Failure default {object} response.Data
// @Router /asset/ [get]
func (h *Handler) getAssets(c *gin.Context) {
	var input assetDto.GetAssetsDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	assets, count, err := h.services.Asset.GetAssets(c, input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: assets,
		Meta: struct {
			Total       int64 `json:"total"`
			CurrentPage int   `json:"current_page"`
		}{
			Total:       count,
			CurrentPage: input.Page,
		},
	})
}

// @Summary Удаление ассета
// @Tags API для работы с ассетами
// @Description Удаление одного файла по его имени
// @ID asset-delete
// @Accept json
// @Produce json
// @Param Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Param name path string true "Название файла"
// @Success 200 {object} uploadResp "Статус"
// @Failure default {object} response.Data
// @Router /asset/{name} [delete]
func (h *Handler) deleteAsset(c *gin.Context) {
	var input assetDto.DeleteAssetDTO
	if err := h.helpers.BindData(c, &input); err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	name := c.Param("name")
	if len(name) < 3 || len(name) > 255 {
		h.helpers.ErrorsHandle(c, models.ErrBadAssetName)
		return
	}

	input.Name = name

	err := h.services.Asset.DeleteAsset(c, input)
	if err != nil {
		h.helpers.ErrorsHandle(c, err)
		return
	}

	response.JsonResponse(c.Writer, response.Data{
		Data: map[string]string{
			"status": "ok",
		},
	})
}
