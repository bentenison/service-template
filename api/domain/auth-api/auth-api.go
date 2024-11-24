package authapi

import (
	"context"
	"net/http"
	"time"

	"github.com/bentenison/microservice/api/domain/auth-api/grpc/proto"
	"github.com/bentenison/microservice/app/domain/authapp"
	"github.com/bentenison/microservice/app/sdk/apperrors"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	authapp *authapp.App
	log     *logger.CustomLogger
	proto.UnimplementedAuthServiceServer
}

func newAPI(log *logger.CustomLogger, authapp *authapp.App) *api {
	return &api{
		authapp: authapp,
		log:     log,
	}
}

func (a *api) checkHealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "running",
		"time":   time.Now(),
	})
}

func (s *api) createUserHandler(c *gin.Context) {
	var userPayload authapp.UserPayload
	if err := c.Bind(&userPayload); err != nil {
		s.log.Errorc(c.Request.Context(), "error while binding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, apperrors.NewError(err))
		return
	}
	id, err := s.authapp.CreateUser(c.Request.Context(), userPayload)
	if err != nil {
		s.log.Errorc(c.Request.Context(), "error in creating user", map[string]interface{}{
			"error": err.Error(),
		})
		// err := apperrors.NewError(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, id)
}
func (s *api) loginHandler(c *gin.Context) {
	var credentials authapp.Credentials
	if err := c.Bind(&credentials); err != nil {
		s.log.Errorc(c.Request.Context(), "error while binding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, apperrors.NewError(err))
		return
	}
	token, err := s.authapp.Authenticate(c.Request.Context(), credentials)
	if err != nil {
		s.log.Errorc(c.Request.Context(), "error in authenticating", map[string]interface{}{
			"error": err.Error(),
		})
		// err := apperrors.NewError(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
func (s *api) authorizeHandler(c *gin.Context) {
	var token token
	if err := c.Bind(&token); err != nil {
		s.log.Errorc(c.Request.Context(), "error while binding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, apperrors.NewError(err))
		return
	}
	id, err := s.authapp.Authorize(c.Request.Context(), token.Token)
	if err != nil {
		s.log.Errorc(c.Request.Context(), "error in authorizing user", map[string]interface{}{
			"error": err.Error(),
		})
		// err := apperrors.NewError(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, id)
}

func (s *api) Authenticate(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	credentials := protoToCred(req)
	token, err := s.authapp.Authenticate(ctx, credentials)
	if err != nil {
		s.log.Errorc(ctx, "error in authenticating", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	return &proto.LoginResponse{Token: token}, nil
}
func (s *api) CreateUser(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	payload := protoToUser(req)
	id, err := s.authapp.CreateUser(ctx, payload)
	if err != nil {
		s.log.Errorc(ctx, "error in creating user", map[string]interface{}{
			"error":    err.Error(),
			"protocol": "GRPC",
		})
		return nil, err
	}
	return &proto.CreateAccountResponse{
		Id:      id,
		Message: "user created",
	}, nil

}
func (s *api) Authorize(ctx context.Context, req *proto.AuthorizeRequest) (*proto.AuthorizeResponse, error) {
	return nil, nil
}
