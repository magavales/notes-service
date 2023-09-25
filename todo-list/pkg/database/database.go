package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo-list/pkg/database/table"
)

type Database struct {
	Pool   *pgxpool.Pool
	Access table.DataAccess
}

func (db *Database) Connect(conf Config) error {
	poolConn, _ := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_max_conns=%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Name, conf.Conns))

	var err error
	db.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConn)

	return err
}
