package userHandler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mediaStorer/param/paramuser"
	"net/http"
)

func (h Handler) userRegisterHandler(e echo.Context) error {
	req := paramuser.RegisterRequest{}
	e.Bind(&req)
	fmt.Println("ph1111", req.PhoneNumber)
	res, err := h.usersvc.Register(req)
	if err != nil {
		fmt.Println("register error", err)
		return err
	}
	//validate data
	// todo validation
	// transport to service
	return e.JSON(http.StatusOK, echo.Map{
		"message": "ok",
		"user":    res.Id,
	})
}
