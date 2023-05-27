package authentication

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user-id"`
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
