package app

import (
	"github.com/gin-gonic/gin"
	"os"
)

const defaultAddr = ":8080"

type App struct {
	srv  *gin.Engine
	addr string
}

func NewFromEnv() *App {
	srv := gin.Default()
	addr := os.Getenv("SERVICE_URL")
	if addr == "" {
		addr = defaultAddr
	}

	return &App{srv: srv, addr: addr}
}

func (a *App) Run() error {
	return a.srv.Run(a.addr)
}

func (a *App) GET(relativePath string, handler Handler) {
	a.srv.GET(relativePath, a.handlerWrapper(handler))
}

func (a *App) POST(relativePath string, handler Handler) {
	a.srv.POST(relativePath, a.handlerWrapper(handler))
}

func (a *App) DELETE(relativePath string, handler Handler) {
	a.srv.DELETE(relativePath, a.handlerWrapper(handler))
}

func (a *App) PATCH(relativePath string, handler Handler) {
	a.srv.PATCH(relativePath, a.handlerWrapper(handler))
}

func (a *App) handlerWrapper(handler Handler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := &Request{ctx: ctx}
		response, err := handler(ctx, request)

		switch err {
		case nil:
			ctx.JSON(200, gin.H{
				"result": gin.H(response),
			})
		case ValidationError:
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		case NotFound:
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
		default:
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
		}
	}
}
