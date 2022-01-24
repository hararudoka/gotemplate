package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HelloHandler struct {

}

func (h *HelloHandler) Call(g *echo.Group) {
	g.GET("", h.Hello)
}


func (h *HelloHandler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
