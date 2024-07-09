package rest

import (
	"github.com/dnevsky/veryGoodProject/internal/configs"
	"github.com/dnevsky/veryGoodProject/internal/service"
	"github.com/dnevsky/veryGoodProject/transport/rest/helpers"
	"github.com/dnevsky/veryGoodProject/transport/rest/middleware"
	v1 "github.com/dnevsky/veryGoodProject/transport/rest/v1"
	"net/http"

	"github.com/Depado/ginprom"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
	cfg      configs.Config
	helpers  helpers.Helpers
}

func NewHandler(services *service.Services, cfg configs.Config, helperManager *helpers.Manager) *Handler {
	return &Handler{
		services: services,
		cfg:      cfg,
		helpers:  helperManager,
	}
}

func (h *Handler) InitRoutes(cfg configs.Config) *gin.Engine {
	router := gin.New()

	if cfg.Env == configs.ProdEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	prometheus := ginprom.New(
		ginprom.Engine(router),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	router.Use(
		middleware.PanicRecovery(),
		middleware.Limit(cfg.Limiter.RPS, cfg.Limiter.Burst, cfg.Limiter.TTL),
		middleware.Cors(),
		prometheus.Instrument(),
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	if cfg.Debug {
		pprof.Register(router)
	}

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	h.InitRoot(router)
	handlerV1 := v1.NewHandler(h.services, h.cfg, h.helpers)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
