package httpServer

import (
	"github.com/labstack/echo/v4"
	"mediaStorer/delivery/httpServer/userHandler"
	"mediaStorer/service/userservice"
)

type Server struct {
	UserHandler userHandler.Handler
	Echo        *echo.Echo
}

func New(service userservice.Service) Server {
	return Server{
		UserHandler: userHandler.New(service),
		Echo:        echo.New(),
	}
}

func (s Server) Serve() {
	s.UserHandler.SetRoute(s.Echo)
	s.Echo.Start(":8088")
}
