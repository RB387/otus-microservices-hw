package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/RB387/otus-microservices-hw/hw5/service-auth/internal/token"
)

type Handler struct {
	tokenManager token.Manager
}

func New(tokenManager token.Manager) *Handler {
	return &Handler{tokenManager: tokenManager}
}

func (h *Handler) Handle(ctx *gin.Context) {
	tkn := ctx.GetHeader("x-auth-token")
	if tkn == "" {
		ctx.JSON(401, gin.H{
			"error": "auth token required",
		})
		return
	}

	claims, err := h.tokenManager.Verify(tkn)

	if err != nil {
		ctx.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}

	username, ok := claims["username"]
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "bad token",
		})
		return
	}

	strUsername, ok := username.(string)
	if !ok {
		ctx.JSON(400, gin.H{
			"error": "bad token",
		})
		return
	}

	ctx.Header("x-user-id", strUsername)
	ctx.JSON(200, gin.H{
		"result": gin.H{
			"ok": true,
		},
	})
}
