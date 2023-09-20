package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"todo-list/pkg/database/table"
)

type Database struct {
	Pool   *pgxpool.Pool
	Access table.DataAccess
}

func (db *Database) Connect() {
	poolConn, _ := pgxpool.ParseConfig("user=postgres password=1703 host=localhost port=5432 dbname=postgres pool_max_conns=10")

	var err error
	db.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConn)
	if err != nil {
		log.Printf("I can't connect to database: %s\n", err)
	}
}
