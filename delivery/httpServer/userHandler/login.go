package userHandler

import (
	"github.com/labstack/echo/v4"
	"mediaStorer/param/paramuser"
	"net/http"
)

func (h Handler) userLoginHandler(e echo.Context) error {
	//bind data
	user := paramuser.LoginRequest{}
	e.Bind(&user)
	// todo validate data

	//pass to service
	userlogin, err := h.usersvc.Login(e.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//return
	return e.JSON(http.StatusOK, userlogin)
}
