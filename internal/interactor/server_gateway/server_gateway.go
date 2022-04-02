package server_gateway

import (
	"github.com/higuruchi/bulletin-board-sample.git/internal/message"
)

type Server interface {
	Run() error
}

// echo.goなどから見せるためのインタフェース
type Controller interface {
	Hello() string
	Post(message.Message) error
	GetAllMessage() ([]message.Message, error)
}