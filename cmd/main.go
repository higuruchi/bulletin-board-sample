package main

import (
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor"
	"github.com/higuruchi/bulletin-board-sample.git/internal/server"
)

const (
	PORT = 1323
)

func main() {
	controller := interactor.NewController()
	server := server.New(PORT, controller)
	interactor := interactor.NewInteractor(server)

	err := interactor.Run()
	if err != nil {
		panic(err)
	}
}