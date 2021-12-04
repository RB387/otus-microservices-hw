package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Request struct {
	ctx    *gin.Context
	claims *jwt.MapClaims
}

func (r *Request) UserID() string {
	return r.ctx.GetHeader("x-user-id")
}

func (r *Request) ParseJSON(obj interface{}) error {
	return r.ctx.ShouldBindJSON(obj)
}

func (r *Request) Param(key string) string {
	return r.ctx.Param(key)
}

func (r *Request) Query(key string) string {
	return r.ctx.Query(key)
}

type Response map[string]interface{}

type Handler func(ctx context.Context, r *Request) (Response, error)
