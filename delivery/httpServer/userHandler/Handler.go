package userHandler

import "mediaStorer/service/userservice"

type Handler struct {
	usersvc userservice.Service
}

func New(service userservice.Service) Handler {
	return Handler{
		service,
	}
}
