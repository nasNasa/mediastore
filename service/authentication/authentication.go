package authentication

import (
	"github.com/golang-jwt/jwt/v4"
	"mediaStorer/entity/userEntity"
	"time"
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
	tokenString, err := token.SignedString(s.config.SignKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
