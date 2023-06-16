package userHandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoute(echo *echo.Echo) {
	userGroup := echo.Group("/user")
	userGroup.POST("/register", h.userRegisterHandler)
	userGroup.POST("/login", h.userLoginHandler)
}
