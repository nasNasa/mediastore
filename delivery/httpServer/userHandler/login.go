package userHandler

import (
	"fmt"
	"mediaStorer/param/paramuser"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) userLoginHandler(e echo.Context) error {
	//bind data
	user := paramuser.LoginRequest{}
	e.Bind(&user)
	// todo validate data

	//pass to service
	userlogin, err := h.usersvc.Login(e.Request().Context(), user)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//return
	return e.JSON(http.StatusOK, userlogin)
}
