package authentication

import (
	"fmt"
	"mediaStorer/entity/userEntity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	config Config
}

func New(config Config) Service {
	return Service{config: config}
}

type Config struct {
	SignKey               string
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string
}

func (s Service) CreateAccessToken(user userEntity.User) (string, error) {
	return s.CreateToken(user.Id, s.config.AccessSubject, s.config.AccessExpirationTime)
}
func (s Service) CreateRefreshToken(user userEntity.User) (string, error) {
	return s.CreateToken(user.Id, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) CreateToken(userId uint, subject string, expiretionTime time.Duration) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiretionTime)),
		},
		UserId: userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.config.SignKey))
	fmt.Println(s.config.SignKey)
	if err != nil {
		fmt.Println("here motherfucker", err)
		return "", err
	}

	return tokenString, nil
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {
	//https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-ParseWithClaims-CustomClaimsType

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
