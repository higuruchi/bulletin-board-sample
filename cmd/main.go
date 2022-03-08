package main

import (
	"github.com/higuruchi/bulletin-board-sample.git/internal/database"
	"github.com/higuruchi/bulletin-board-sample.git/internal/db_handler"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor"
	"github.com/higuruchi/bulletin-board-sample.git/internal/server"
)

const (
	Port = 1323
	TextPath = "./messages.txt"
)

func main() {
	textDB, err := database.NewText(TextPath)
	if err != nil {
		panic(err)
	}
	DBHandler := db_handler.New(textDB)
	controller := interactor.NewController(DBHandler)
	server := server.New(Port, controller)
	interactor := interactor.NewInteractor(server)

	err = interactor.Run()
	if err != nil {
		panic(err)
	}
}