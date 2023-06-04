package paramuser

import "mediaStorer/entity/userEntity"

type LoginRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type LoginResponse struct {
	User         userEntity.User
	AccessToken  string
	RefreshToken string
}
