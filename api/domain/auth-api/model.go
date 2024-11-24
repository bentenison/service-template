package authapi

import (
	"github.com/bentenison/microservice/api/domain/auth-api/grpc/proto"
	"github.com/bentenison/microservice/app/domain/authapp"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type token struct {
	Token string `json:"token,omitempty"`
}

func protoToUser(req *proto.CreateAccountRequest) authapp.UserPayload {
	var u authapp.UserPayload
	// u.ID = uuid.NewString()
	// u.CreatedAt = time.Now()
	// u.UpdatedAt = time.Now()
	u.Username = req.Username
	u.Password = req.Password
	u.Email = req.Email
	u.Role = req.Role
	return u
}
func protoToCred(req *proto.LoginRequest) authapp.Credentials {
	var u authapp.Credentials
	u.Username = req.Username
	u.Password = req.Password
	return u
}
