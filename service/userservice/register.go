package userservice

import (
	"fmt"
	"mediaStorer/entity/userEntity"
	"mediaStorer/param/paramuser"
)

func (s Service) Register(request paramuser.RegisterRequest) (userEntity.User, error) {

	//hash pass

	//save to db
	fmt.Println("request ph", request.PhoneNumber)
	user := userEntity.User{
		Email:       request.Email,
		Name:        request.Name,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
	}
	user, err := s.repository.RegisterToDb(user)
	if err != nil {
		return userEntity.User{}, fmt.Errorf("registering in datanase error:%s", err)
	}
	return user, nil
}
