package authapp

import (
	"context"

	"github.com/bentenison/microservice/business/domain/authbus"
	"github.com/bentenison/microservice/foundation/conf"
	"github.com/bentenison/microservice/foundation/logger"
	tp "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	authbus   *authbus.Business
	logger    *logger.CustomLogger
	tracer    trace.Tracer
	appConfig *conf.Config
}

func NewApp(log *logger.CustomLogger, authbus *authbus.Business, tp *tp.TracerProvider, conf *conf.Config) *App {
	return &App{
		authbus:   authbus,
		logger:    log,
		tracer:    tp.Tracer("AUTH"),
		appConfig: conf,
	}
}

func (a *App) CreateUser(ctx context.Context, u UserPayload) (string, error) {
	user := toBusUser(&u)
	id, err := a.authbus.CreateUser(ctx, user)
	if err != nil {
		a.logger.Errorc(ctx, "error in creating user", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	return id, nil
}
func (a *App) Authenticate(ctx context.Context, cred Credentials) (string, error) {
	token, err := a.authbus.AuthenticateUser(ctx, cred.Username, cred.Password, a.appConfig.JWTKey)
	if err != nil {
		a.logger.Errorc(ctx, "error while creating user token", map[string]interface{}{
			"error": err.Error(),
		})
		return "", err
	}
	return token, nil
}
func (a *App) Authorize(ctx context.Context, token string) (*authbus.Claims, error) {

	claims, err := a.authbus.AuthorizeUser(ctx, token, a.appConfig.JWTKey)
	if err != nil {
		a.logger.Errorc(ctx, "error while authorizing user token", map[string]interface{}{
			"error": err.Error(),
		})
		return &authbus.Claims{}, err
	}

	return claims, nil
}
