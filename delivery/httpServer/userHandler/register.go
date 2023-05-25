package userHandler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mediaStorer/service/userservice"
	"net/http"
)

func (h Handler) registerHandler(e echo.Context) error {
	req := userservice.RegisterRequest{}
	e.Bind(&req)
	fmt.Println(req.Email)
	res, err := h.usersvc.Register(req)
	if err != nil {
		fmt.Println("register error", err)
	}
	//validate data

	// transport to service
	return e.JSON(http.StatusOK, echo.Map{
		"message": "ok",
		"user":    res.Id,
	})
}
