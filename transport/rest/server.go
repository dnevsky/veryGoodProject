package rest

import (
	"context"
	"github.com/dnevsky/veryGoodProject/internal/configs"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunHttp(cfg configs.Config, handler http.Handler) error {
	var writeTimeout time.Duration
	if cfg.Debug {
		writeTimeout = time.Second * 120
	} else {
		writeTimeout = cfg.HTTPConfig.WriteTimeout
	}

	// certFile := cfg.HTTPSConfig.CertFile
	// keyFile := cfg.HTTPSConfig.KeyFile
	//
	// cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// if err != nil {
	// 	return err
	// }
	//
	// tlsConfig := &tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// }

	logger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	s.httpServer = &http.Server{
		Addr:              ":" + cfg.HTTPConfig.Port,
		Handler:           handler,
		ReadTimeout:       cfg.HTTPConfig.ReadTimeout,
		WriteTimeout:      writeTimeout,
		MaxHeaderBytes:    cfg.HTTPConfig.MaxHeaderMegabytes << 20,
		ReadHeaderTimeout: cfg.HTTPConfig.ReadTimeout,
		IdleTimeout:       cfg.HTTPConfig.ReadTimeout,
		ErrorLog:          logger,
		// TLSConfig:         tlsConfig,
	}

	return s.httpServer.ListenAndServe()
	// return s.httpsServer.ListenAndServeTLS("", "")
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
