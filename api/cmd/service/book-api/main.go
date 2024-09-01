package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bentenison/microservice/api/cmd/service/book-api/build/all"
	"github.com/bentenison/microservice/api/sdk/http/debug"
	"github.com/bentenison/microservice/api/sdk/http/mux"
	"github.com/bentenison/microservice/business/sdk/sqldb"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

var apiType = "all"

func main() {
	//initialize the application logger
	log := logger.NewCustomLogger(map[string]interface{}{
		"service": "example-service",
		"env":     "production",
		"build":   "1.0.0",
	})
	// load congigurations
	config, err := conf.LoadConfig()
	if err != nil {
		log.Error("error while loading conf", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// // -------------------------------------------------------------------------
	// // INITIALIZE TRACER OTEL
	trace, err := otel.NewTracer()
	if err != nil {
		log.Error("error while initializing tracer", map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer func() {
		otel.ShutDownTracer(trace)
	}()
	log.Info("config", map[string]interface{}{"config": config})
	if err := run(log, trace, config); err != nil {
		log.Error("error while running server", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// log.Error("error while loading conf", map[string]interface{}{
	// 	"error": "error",
	// })
}
func run(log *logger.CustomLogger, tracer *trace.TracerProvider, cfg *conf.Config) error {
	//starting database connection
	db, err := sqldb.Open(sqldb.Config{
		User:         cfg.User,
		Password:     cfg.Password,
		Host:         cfg.Host,
		Name:         cfg.DBName,
		MaxIdleConns: cfg.MaxIdleConns,
		MaxOpenConns: cfg.MaxOpenConns,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}

	defer db.Close()
	go func() {
		log.Info("startup debug v1 server started", map[string]interface{}{
			"port": cfg.DebugPort,
		})

		if err := http.ListenAndServe(cfg.DebugPort, debug.Mux()); err != nil {
			log.Error("error occured while listning for traffic", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := mux.Config{
		Build:  "develop",
		Log:    log,
		DB:     db,
		Tracer: tracer,
	}

	api := http.Server{
		Addr:    cfg.BookAPIPort,
		Handler: mux.WebAPI(cfgMux, buildRoutes()),
		// ReadTimeout:  cfg.Web.ReadTimeout,
		// WriteTimeout: cfg.,
		// IdleTimeout:  cfg.Web.IdleTimeout,
		// ErrorLog: lo,
	}

	serverErrors := make(chan error, 1)
	ctx := context.Background()
	go func() {
		log.Info("api router started", map[string]interface{}{
			"port": cfg.BookAPIPort,
		})
		serverErrors <- api.ListenAndServe()
	}()

	// // -------------------------------------------------------------------------
	// // Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info("shutdown started", map[string]interface{}{
			"signal": sig,
		})
		defer log.Info("shutdown completed")

		ctx, cancel := context.WithTimeout(ctx, time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}

func buildRoutes() mux.RouteAdder {
	switch apiType {
	case "all":
		return all.Routes()

	}
	return all.Routes()
}
