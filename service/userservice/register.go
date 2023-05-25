package userservice

import (
	"fmt"
	"mediaStorer/entity/userEntity"
)

type RegisterRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone-number"`
}

func (s Service) Register(request RegisterRequest) (userEntity.User, error) {

	//hash pass

	//save to db
	user := userEntity.User{
		Email:       request.Email,
		Name:        request.Name,
		Password:    request.Password,
		PhoneNumber: request.Password,
	}
	user, err := s.repository.RegisterToDb(user)
	if err != nil {
		return userEntity.User{}, fmt.Errorf("registering in datanase error:%s", err)
	}
	return user, nil
}
