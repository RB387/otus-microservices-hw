package register

import (
	"github.com/gin-gonic/gin"

	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/clients/user"
	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/token"
)

type Handler struct {
	userClient   user.Client
	tokenManager token.Manager
}

func New(users user.Client, tokenManager token.Manager) *Handler {
	return &Handler{userClient: users, tokenManager: tokenManager}
}

func (h *Handler) Handle(ctx *gin.Context) {
	req := &User{}
	err := ctx.BindJSON(req)
	if err != nil {
		return
	}

	err = h.userClient.Create(user.User{
		Username:  req.Username,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	})

	switch err {
	case nil:
		tkn, _ := h.tokenManager.Sign(req.Username)
		ctx.Header("x-auth-token", tkn)
		ctx.Header("x-user-id", req.Username)
		ctx.JSON(200, gin.H{
			"result": gin.H{
				"x-auth-token": tkn,
			},
		})
	default:
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}
