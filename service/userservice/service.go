package userservice

import "mediaStorer/entity/userEntity"

type Service struct {
	repository Repository
}

type Repository interface {
	RegisterToDb(request userEntity.User) (userEntity.User, error)
}

func New(repository Repository) Service {
	return Service{repository: repository}
}
