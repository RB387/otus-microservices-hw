package app

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Request struct {
	ctx *gin.Context
}

func (r *Request) ParseJSON(obj interface{}) error {
	return r.ctx.ShouldBindJSON(obj)
}

func (r Request) Param(key string) string {
	return r.ctx.Param(key)
}

type Response map[string]interface{}

type Handler func(ctx context.Context, r *Request) (Response, error)
