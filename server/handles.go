package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) getUserInfo(ctx echo.Context) error {
	c := ctx.(*CustomContext)

	assets := "some"
	return ResponseSuccess(c, assets)
}

func (s *Server) index(ctx echo.Context) error {
	c := ctx.(*CustomContext)
	return c.Render(http.StatusOK, "index.html", nil)
}

func (s *Server) show(ctx echo.Context) error {
	c := ctx.(*CustomContext)
	return c.Render(http.StatusOK, "show.html", nil)
}
