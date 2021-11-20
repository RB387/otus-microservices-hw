package health

import (
	"context"
	"github.com/RB387/otus-microservices-hw/hw2/service-user/internal/app"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(_ context.Context, _ *app.Request) (app.Response, error) {
	return app.Response{
		"ok": true,
	}, nil
}
