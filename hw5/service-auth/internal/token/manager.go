package token

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const defaultExp = "5m"

type manager struct {
	key []byte
	exp time.Duration
}

type Manager interface {
	Sign(username string) (string, error)
	Verify(token string) (jwt.MapClaims, error)
}

func NewManager() (*manager, error) {
	key := os.Getenv("JWT_TOKEN_KEY")
	if key == "" {
		return nil, errors.New("jwt token key must be set in environ")
	}

	exp := os.Getenv("JWT_TOKEN_EXP")
	if exp == "" {
		exp = defaultExp
	}

	expDuration, err := time.ParseDuration(exp)
	if err != nil {
		return nil, err
	}

	return &manager{key: []byte(key), exp: expDuration}, nil
}
