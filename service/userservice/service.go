package userservice

import (
	"context"
	"mediaStorer/entity/userEntity"
)

type Service struct {
	repository Repository
}

type Repository interface {
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (userEntity.User, error)
	RegisterToDb(request userEntity.User) (userEntity.User, error)
}

func New(repository Repository) Service {
	return Service{repository: repository}
}
