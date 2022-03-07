package interactor

import (
	"fmt"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor/server"
)

type interactor struct {
	server server.Server
}

type controller struct {}

// main.goにから見せるためのインタフェース
type Interactor interface {
	Run() error
}

// echo.goなどから見せるためのインタフェース
type Controller interface {
	Hello() string
}

func NewInteractor(s server.Server) Interactor {
	return &interactor {
		server: s,
	} 
}

func NewController() Controller {
	return &controller{}
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