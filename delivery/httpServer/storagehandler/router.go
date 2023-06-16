package storagehandler

import (
	"mediaStorer/delivery/httpServer/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoute(echo *echo.Echo) {
	userGroup := echo.Group("/file")
	userGroup.POST("/write", h.writeFileHandler,
		middleware.Auth(h.authSvc, h.authCfg))
	userGroup.GET("/download", h.downloadFileHandler,
		middleware.Auth(h.authSvc, h.authCfg))
}
