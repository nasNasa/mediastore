package userservice

import "mediaStorer/entity/userEntity"

type LoginRequest struct {
}

func (s Service) Login(request LoginRequest) (userEntity.User, error) {
	//get user from database

	//verify password

	//create access token and refresh token

	//return entity and token
}
