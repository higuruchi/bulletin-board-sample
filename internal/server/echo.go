package server

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/higuruchi/bulletin-board-sample.git/internal/message"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor/server_gateway"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor"
)

type (
	server struct {
		port int
		echoImplement *echo.Echo
		controller interactor.Controller
	}
	
	Messages struct {
		Len int `json: len`
		Messages []string `json: messages`
	}

	Message struct {
		Message string `json: message`
	}
)

func New(port int, c interactor.Controller) server_gateway.Server {
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

	s.echoImplement.GET("/messages", func(c echo.Context) error {
		messages, err := s.controller.GetAllMessage()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		retMessages := []string{}
		for _, message := range messages {
			retMessages = append(retMessages, message.Message())
		}

		retJSON := Messages {
			Len: len(messages),
			Messages: retMessages,
		}

		return c.JSON(http.StatusOK, retJSON)
	})

	s.echoImplement.POST("/message", func(c echo.Context) error {
		postMessage := new(Message)
		if err := c.Bind(postMessage); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		message := message.New(postMessage.Message)

		err := s.controller.Post(message)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "")
	})

	s.echoImplement.Start(fmt.Sprintf(":%v", s.port))
	return nil
}