package token

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var InvalidToken = errors.New("invalid token")

func (m *manager) Verify(token string) (jwt.MapClaims, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return m.key, nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	validErr := parsed.Claims.Valid()
	if !parsed.Valid || validErr != nil {
		return jwt.MapClaims{}, InvalidToken
	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return jwt.MapClaims{}, InvalidToken
	}
}
