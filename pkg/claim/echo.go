package claim

import (
	"mediaStorer/config"
	"mediaStorer/service/authentication"

	"github.com/labstack/echo/v4"
)

func GetClaimsFromEchoContext(c echo.Context) *authentication.Claims {
	return c.Get(config.AuthMiddlewareContextKey).(*authentication.Claims)
}
