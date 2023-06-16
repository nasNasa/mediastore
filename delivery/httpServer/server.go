package httpServer

import (
	"mediaStorer/delivery/httpServer/storagehandler"
	"mediaStorer/delivery/httpServer/userHandler"
	"mediaStorer/service/authentication"
	"mediaStorer/service/fileservice"
	"mediaStorer/service/userservice"

	"github.com/labstack/echo/v4"
)

type Server struct {
	AuthConfig     authentication.Config
	StorageHandler storagehandler.Handler
	UserHandler    userHandler.Handler
	Echo           *echo.Echo
}

func New(AuthConfig authentication.Config, service userservice.Service, service2 authentication.Service, service3 fileservice.Service) Server {
	return Server{
		StorageHandler: storagehandler.New(AuthConfig, service, service2, service3),
		UserHandler:    userHandler.New(service, service2),
		Echo:           echo.New(),
	}
}

func (s Server) Serve() {
	s.StorageHandler.SetRoute(s.Echo)
	s.UserHandler.SetRoute(s.Echo)
	s.Echo.Start(":8088")
}
