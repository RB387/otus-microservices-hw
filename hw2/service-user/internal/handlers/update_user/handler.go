package update_user

import (
	"context"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/app"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/storages/user"
)

type Storage interface {
	Update(ctx context.Context, user user.User) error
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

	var usr User

	err := r.ParseJSON(&usr)
	if err != nil {
		return nil, app.ValidationError
	}

	err = h.storage.Update(ctx, user.User{
		Username:  userName,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Phone:     usr.Phone,
	})

	switch err {
	case user.NotFound:
		return nil, app.NotFound
	case nil:
		return map[string]interface{}{
			"ok": true,
		}, nil
	default:
		return nil, err
	}
}
