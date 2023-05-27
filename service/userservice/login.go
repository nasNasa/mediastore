package userservice

import (
	"context"
	"fmt"
	"mediaStorer/entity/userEntity"
	"mediaStorer/param/paramuser"
)

func (s Service) Login(ctx context.Context, request paramuser.LoginRequest) (userEntity.User, error) {
	//get user from database
	fmt.Println("request", request)
	user, err := s.repository.GetUserByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return userEntity.User{}, err
	}
	//verify password

	//create access token and refresh token

	//return entity and token
	return user, nil
}
