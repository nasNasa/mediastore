package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"mediaStorer/entity/userEntity"
)

type Service struct {
	repository Repository
	auth       AuthGenerator
}

type Repository interface {
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userEntity.User, error)
	RegisterToDb(request userEntity.User) (userEntity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user userEntity.User) (string, error)
	CreateRefreshToken(user userEntity.User) (string, error)
}

func New(repository Repository, authG AuthGenerator) Service {
	return Service{repository: repository, auth: authG}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
