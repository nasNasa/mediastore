package userservice

import (
	"context"
	"fmt"
	"mediaStorer/param/paramuser"
)

func (s Service) Login(ctx context.Context, request paramuser.LoginRequest) (paramuser.LoginResponse, error) {
	//get user from database
	fmt.Println("request", request)
	user, err := s.repository.GetUserByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		return paramuser.LoginResponse{}, err
	}
	//verify password
	if user.Password != getMD5Hash(request.Password) {
		return paramuser.LoginResponse{}, fmt.Errorf("password or phone number is incorrect")
	}
	//create access token and refresh token
	accessToken, cErr := s.auth.CreateAccessToken(user)
	if cErr != nil {
		return paramuser.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	refreshToken, rErr := s.auth.CreateAccessToken(user)
	if rErr != nil {
		return paramuser.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	//return entity and token
	return paramuser.LoginResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
