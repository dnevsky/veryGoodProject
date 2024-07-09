package app

import (
	"context"
	"errors"
	"github.com/dnevsky/veryGoodProject/internal/configs"
	"github.com/dnevsky/veryGoodProject/internal/repository"
	"github.com/dnevsky/veryGoodProject/internal/repository/postgresDB"
	"github.com/dnevsky/veryGoodProject/internal/service"
	"github.com/dnevsky/veryGoodProject/pkg/logger"
	"github.com/dnevsky/veryGoodProject/transport/rest"
	"github.com/dnevsky/veryGoodProject/transport/rest/helpers"
	_ "github.com/lib/pq"
	"github.com/mitchellh/panicwrap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	cfg, err := configs.Init(configPath)
	if err != nil {
		log.Panicln(err)
		return
	}

	exitStatus, err := panicwrap.BasicWrap(panicHandler)
	if err != nil {
		log.Fatal(err)
	}
	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}

	db, err := postgresDB.NewPostgresDB(cfg.DB.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer postgresDB.Close(db)

	repositories := repository.NewRepositories(db)

	logManager := logger.NewLogManager()

	services, err := service.NewServices(service.Deps{
		Repositories: repositories,
		Logger:       logManager,
		Config:       cfg,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	helpersManager := helpers.NewManager(
		logManager,
	)

	httpServerInstance := new(rest.Server)

	handlers := rest.NewHandler(services, cfg, helpersManager)

	go func() {
		if err := httpServerInstance.RunHttp(cfg, handlers.InitRoutes(cfg)); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	logManager.Info("Started..")

	<-quit

	if err := httpServerInstance.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occurated on shuting down server: %s", err.Error())
	}

}

func panicHandler(output string) {
	// тут можно ошибку куда-нибудь залогировать
	os.Exit(1)
}
