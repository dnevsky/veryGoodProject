package v1

import (
	"github.com/dnevsky/veryGoodProject/docs"
	"github.com/dnevsky/veryGoodProject/internal/configs"
	"github.com/dnevsky/veryGoodProject/internal/service"
	"github.com/dnevsky/veryGoodProject/transport/rest/helpers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
	config   configs.Config
	helpers  helpers.Helpers
}

func NewHandler(services *service.Services, cfg configs.Config, helpers helpers.Helpers) *Handler {
	return &Handler{
		services: services,
		config:   cfg,
		helpers:  helpers,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initAssetRoutes(v1)
	}
}
