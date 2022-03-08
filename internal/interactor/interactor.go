package interactor

import (
	"fmt"
	"errors"
	"github.com/higuruchi/bulletin-board-sample.git/internal/message"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor/server_gateway"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor/db_gateway"
)

var (
	MessageNotFound = errors.New("Errors Not Found")
)

type interactor struct {
	server server_gateway.Server
}

type controller struct {
	database db_gateway.DB
}

// main.goにから見せるためのインタフェース
type Interactor interface {
	Run() error
}

// echo.goなどから見せるためのインタフェース
type Controller interface {
	Hello() string
	Post(message.Message) error
	GetAllMessage() ([]message.Message, error)
}

func NewInteractor(s server_gateway.Server) Interactor {
	return &interactor {
		server: s,
	} 
}

func NewController(d db_gateway.DB) Controller {
	return &controller{
		database: d,
	}
}

func (i *interactor)Run() error {
	err := i.server.Run()
	if err != nil {
		return fmt.Errorf("Calling Interactor.Run: %w", err)
	}

	return nil
}

func (c *controller)Hello() string {
	return "Hello"
}

func (c *controller)Post(m message.Message) error {
	if len(m.Message()) == 0 {
		return MessageNotFound
	}

	err := c.database.RecordMessage(m)
	if err != nil {
		return fmt.Errorf("Calling interactor.Post: %w", err)
	}

	return nil
}

func (c *controller)GetAllMessage() ([]message.Message, error) {
	messages, err := c.database.GetAllMessage()
	if err != nil {
		return nil, fmt.Errorf("Calling interactor.GetAllMessage: %w", err)
	}

	return messages, nil
}