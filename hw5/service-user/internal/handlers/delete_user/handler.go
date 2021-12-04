package delete_user

import (
	"context"
	"fmt"

	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/app"
	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/storages/user"
)

type Storage interface {
	Delete(ctx context.Context, username string) error
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

	if r.UserID() != userName {
		fmt.Printf("denied for %s. trying to get %s", r.UserID(), userName)
		return nil, app.PermissionDenied
	}

	err := h.storage.Delete(ctx, userName)

	switch err {
	case user.NotFound:
		return nil, app.NotFound
	case nil:
		return app.Response{
			"ok": true,
		}, nil
	default:
		return nil, err
	}
}
