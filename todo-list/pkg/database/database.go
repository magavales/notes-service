package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"log"
	"todo-list/pkg/database/table"
)

type Database struct {
	Pool   *pgxpool.Pool
	Access table.DataAccess
}

func (db *Database) Connect() error {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err)
	}
	poolConn, _ := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_max_conns=%s", viper.GetString("user"), viper.GetString("password"), viper.GetString("host"), viper.GetString("port"), viper.GetString("dbname"), viper.GetString("conns")))

	var err error
	db.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConn)

	return err
}

func initConfig() error {
	viper.SetConfigName("config_database")
	viper.SetConfigType("yml")
	viper.AddConfigPath("configs")

	return viper.ReadInConfig()
}
