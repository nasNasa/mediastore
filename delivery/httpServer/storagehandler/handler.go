package storagehandler

import (
	"mediaStorer/service/authentication"
	"mediaStorer/service/fileservice"
	"mediaStorer/service/userservice"
)

type Handler struct {
	authCfg authentication.Config
	usersvc userservice.Service
	authSvc authentication.Service
	fileSvc fileservice.Service
}

func New(authCfg authentication.Config, service userservice.Service, authService authentication.Service, fileSvc fileservice.Service) Handler {
	return Handler{
		authCfg,
		service,
		authService,
		fileSvc,
	}
}
