package main

import (
	"github.com/higuruchi/bulletin-board-sample.git/internal/database"
	"github.com/higuruchi/bulletin-board-sample.git/internal/db_handler"
	"github.com/higuruchi/bulletin-board-sample.git/internal/interactor"
	"github.com/higuruchi/bulletin-board-sample.git/internal/server"
)

const (
	db_ip = "172.22.2.71"
	db_port = "3306"
	db_user = "user"
	db_password = "password"
	db_name = "bulletin_board"
	port = 1323
)

func main() {
	mysqlDB, f, err := database.NewMySQL(db_user, db_password, db_ip, db_port, db_name)
	if err != nil {
		panic(err)
	}
	defer f()
	DBHandler := db_handler.NewMySQL(mysqlDB)
	controller := interactor.NewController(DBHandler)
	server := server.New(port, controller)
	interactor := interactor.NewInteractor(server)

	e := interactor.Run()
	if e != nil {
		panic(err)
	}
}