package login

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
	req := &AuthInfo{}
	err := ctx.BindJSON(req)
	if err != nil {
		return
	}

	ok, err := h.userClient.CheckPassword(req.Username, req.Password)
	if err == user.UserNotFound {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	} else if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if ok {
		tkn, _ := h.tokenManager.Sign(req.Username)
		ctx.Header("x-auth-token", tkn)
		ctx.Header("x-user-id", req.Username)
		ctx.JSON(200, gin.H{
			"result": gin.H{
				"x-auth-token": tkn,
			},
		})
	} else {
		ctx.JSON(401, gin.H{
			"error": "wrong password",
		})
	}
}
