package main

import (
	"github.com/spf13/viper"
	"log"
	"todo-list/pkg/handler"
	"todo-list/pkg/server"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err)
	}
	router := new(handler.Handler).InitRouter()

	serv := new(server.Server)
	err := serv.InitServer(viper.GetString("port"), router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
func initConfig() error {
	viper.SetConfigName("config_server")
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")

	return viper.ReadInConfig()
}
