package user

import (
	"context"

	"github.com/arfan21/project-sprint-banking-api/internal/model"
)

type Service interface {
	Register(ctx context.Context, req model.UserRegisterRequest) (res model.UserLoginResponse, err error)
	Login(ctx context.Context, req model.UserLoginRequest) (res model.UserLoginResponse, err error)
}
