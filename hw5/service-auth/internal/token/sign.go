package token

import (
	"github.com/golang-jwt/jwt/v4"
)

func (m *manager) Sign(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      jwt.TimeFunc().Add(m.exp).Unix(),
	})

	return token.SignedString(m.key)
}
