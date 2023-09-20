package table

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"todo-list/pkg/model"
)

type DataAccess struct {
}

func (da DataAccess) CreateTask(pool *pgxpool.Pool, task model.Task) error {
	rows, err := pool.Query(context.Background(), "INSERT INTO tasks (header, description, date, status) VALUES ($1, $2, $3, $4) RETURNING id", task.Header, task.Description, task.Date, task.Status)
	if err != nil {
		return err
	}
	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("error while iterating dataset")
		}
		task.ID = values[0].(int64)
	}

	return err
}
