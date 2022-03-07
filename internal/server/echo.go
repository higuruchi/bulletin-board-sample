package server

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"

	server_interface "github.com/higuruchi/bulletin-board-sample.git/internal/interactor/server"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor"
)

type server struct {
	port int
	echoImplement *echo.Echo
	controller interactor.Controller
}

func New(port int, c interactor.Controller) server_interface.Server {
	e := echo.New()

	return &server {
		port: port,
		echoImplement: e,
		controller: c,
	}
}

func (s *server)Run() error {
	s.echoImplement.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, s.controller.Hello())
	})

	s.echoImplement.Start(fmt.Sprintf(":%v", s.port))
	return nil
}