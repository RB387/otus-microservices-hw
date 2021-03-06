package create_user

import (
	"context"

	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/app"
	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/password"
	"github.com/RB387/otus-microservices-hw/hw5/service-user/internal/storages/user"
)

type Storage interface {
	Create(ctx context.Context, user user.User) error
}

type Handler struct {
	storage Storage
}

func New(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) Handle(ctx context.Context, r *app.Request) (app.Response, error) {
	var usr User

	err := r.ParseJSON(&usr)
	if err != nil {
		return nil, app.ValidationError
	}

	salt, err := password.GenerateSalt()
	if err != nil {
		return nil, err
	}

	err = h.storage.Create(ctx, user.User{
		Username:  usr.Username,
		Password:  password.CalculatePassword(usr.Password, salt),
		Salt:      salt,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Phone:     usr.Phone,
	})

	if err != nil {
		return nil, err
	}

	return app.Response{
		"ok": true,
	}, nil
}
