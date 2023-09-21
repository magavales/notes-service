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

func (da DataAccess) GetTaskByID(pool *pgxpool.Pool, id int64) (task model.Task, err error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM tasks WHERE id = $1", id)
	if err != nil {
		return model.Task{}, err
	}

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println("error while iterating dataset")
		}
		task.ParseRowsFromTable(values)
	}

	return task, err
}

func (da DataAccess) UpdateTask(pool *pgxpool.Pool, task model.Task) error {
	var err error
	_, err = pool.Exec(context.Background(), "UPDATE tasks SET header = $2, description = $3, date = $4, status = $5 WHERE id = $1", task.ID, task.Header, task.Description, task.Date.Time, task.Status)
	if err != nil {
		return err
	}

	return err
}

func (da DataAccess) DeleteTask(pool *pgxpool.Pool, id int64) (err error) {
	_, err = pool.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	return err
}
