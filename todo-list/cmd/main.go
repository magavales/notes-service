package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"todo-list/pkg/database"
	"todo-list/pkg/handler"
	"todo-list/pkg/server"
)

/*
@title Todo List API
@version v1.0.0
@description API server for TodoList Application
@host localhost:8080
@BasePath /
@in header
*/
func main() {
	h := new(handler.Handler)
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err)
	}
	h.Config = database.Config{
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Name:     viper.GetString("db.dbname"),
		Conns:    viper.GetString("db.conns"),
	}
	if len(os.Args) == 2 {
		h.Config.Host = os.Args[1]
	}

	router := h.InitRouter()

	serv := new(server.Server)
	err := serv.InitServer(viper.GetString("server.port"), router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
