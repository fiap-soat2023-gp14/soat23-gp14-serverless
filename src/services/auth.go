package services

import (
	"context"

	"github.com/soat-2023-gp14/soat23-gp14_fiap-lambda-application/models"
)

type Auth interface {
	SignUp(context.Context, models.UserForm) error
	Login(context.Context, models.UserLogin) (string, error)
}
