package config

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler struct {
	Repo Repository
}

func (h *Handler) GetAppConfig(c echo.Context) error {
	ac := h.Repo.GetConfig()

	return c.JSON(http.StatusOK, ac.AppConfigData)
}
