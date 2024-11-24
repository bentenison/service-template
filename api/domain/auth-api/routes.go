package authapi

import (
	"context"

	"github.com/bentenison/microservice/api/domain/auth-api/grpc/proto"
	"github.com/bentenison/microservice/api/sdk/grpc/rpcserver"
	"github.com/bentenison/microservice/app/domain/authapp"
	"github.com/bentenison/microservice/app/sdk/mid"
	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/bentenison/microservice/foundation/web"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Log *logger.CustomLogger
	// authapp
	AuthBus   *authbus.Business
	Tracer    *trace.TracerProvider
	AppConfig *conf.Config
}

func Routes(app *web.App, cfg Config) {
	api := newAPI(cfg.Log, authapp.NewApp(cfg.Log, cfg.AuthBus, cfg.Tracer, cfg.AppConfig))
	app.Use(mid.TraceIdMiddleware())
	go RunGRPCServer(cfg.AppConfig.GRPCPort, cfg.Log, api)
	app.Handle("GET", "/auth/check", api.checkHealthHandler)
	app.Handle("POST", "/auth/create", api.createUserHandler)
	app.Handle("POST", "/auth/authenticate", api.loginHandler)
	app.Handle("POST", "/auth/authorize", api.authorizeHandler)
}
func RunGRPCServer(GRPCPort string, log *logger.CustomLogger, api *api) {
	grpcSrv, listner := rpcserver.CreateServer(GRPCPort, log)
	defer listner.Close()
	// go func() {
	log.Infoc(context.TODO(), "startup grpc v1 server started", map[string]interface{}{
		"port": GRPCPort,
	})
	// executorServer := executorapi.NewExecutorServer(log)
	proto.RegisterAuthServiceServer(grpcSrv, api)
	reflection.Register(grpcSrv)
	if err := grpcSrv.Serve(listner); err != nil {
		log.Errorc(context.TODO(), "error occured while listning for grpc traffic", map[string]interface{}{
			"error": err.Error(),
		})
	}
	// }()
}
