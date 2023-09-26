package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	Conns    string
}

func (c *Config) Connect() (*pgxpool.Pool, error) {
	poolConn, err := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_max_conns=%s", c.User, c.Password, c.Host, c.Port, c.Name, c.Conns))
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConn)
	if err != nil {
		return nil, err
	}
	return pool, err
}
