package check_password

import (
	"context"

	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/app"
	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/password"
	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/storages/user"
)

type Storage interface {
	Get(ctx context.Context, username string) (user.User, error)
}

type Handler struct {
	storage   Storage
	paramName string
}

func New(storage Storage, paramName string) *Handler {
	return &Handler{storage: storage, paramName: paramName}
}

func (h *Handler) Handle(ctx context.Context, r *app.Request) (app.Response, error) {
	userName := r.Param(h.paramName)
	if userName == "" {
		return nil, app.ValidationError
	}

	validatePassword := r.Query("password")
	if validatePassword == "" {
		return nil, app.ValidationError
	}

	usr, err := h.storage.Get(ctx, userName)

	switch err {
	case user.NotFound:
		return nil, app.NotFound
	case nil:
		passwordHash := password.CalculatePassword(validatePassword, usr.Salt)
		return app.Response{
			"valid": passwordHash == usr.Password,
		}, nil
	default:
		return nil, err
	}
}
