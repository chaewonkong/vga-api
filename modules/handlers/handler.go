package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Handler struct {
	e *echo.Echo
}

func NewHandler(e *echo.Echo) *Handler {
	return &Handler{e}
}

func BindRoutes(h *Handler) {
	h.e.GET("/alive", h.Alive)
}

func (h Handler) Alive(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
