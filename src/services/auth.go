package services

import (
	"context"

	"github.com/oivinig/soat23-gp14-serverless/models"
)

type Auth interface {
	SignUp(context.Context, models.UserForm) error
	Login(context.Context, models.UserLogin) (string, error)
}
