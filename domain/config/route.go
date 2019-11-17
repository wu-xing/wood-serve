package config

import (
	"github.com/labstack/echo"
	"github.com/wu-xing/wood-serve/common"
)

func RegisterHandler(g *echo.Group) {
	handler := &Handler{Repo: NewRepository(common.GetDB().Connection)}

	g.GET("/config", handler.GetAppConfig)
}
