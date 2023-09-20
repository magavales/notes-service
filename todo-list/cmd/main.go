package main

import (
	"log"
	"todo-list/pkg/handler"
	"todo-list/pkg/server"
)

func main() {
	router := new(handler.Handler).InitRouter()

	serv := new(server.Server)
	err := serv.InitServer("8080", router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
