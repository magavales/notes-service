package table

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"todo-list/pkg/model"
)

type DataAccess struct {
}

func (da DataAccess) CreateTask(pool *pgxpool.Pool, task *model.Task) error {
	rows, err := pool.Query(context.Background(), "INSERT INTO tasks (header, description, date, status) VALUES ($1, $2, $3, $4) RETURNING id", task.Header, task.Description, task.Date.Time, task.Status)
	if err != nil {
		return err
	}
	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Printf("error while iterating dataset. Error: %s\n", err)
			return err
		}
		task.ID = values[0].(int64)
	} else {
		return pgx.ErrNoRows
	}

	log.Printf("Task has been created!")

	return err
}

func (da DataAccess) GetTasks(pool *pgxpool.Pool) (tasks []model.Task, err error) {
	var temp model.Task
	rows, err := pool.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Printf("error while iterating dataset. Error: %s\n", err)
			return tasks, err
		}
		temp.ParseRowsFromTable(values)
		tasks = append(tasks, temp)
	}

	return tasks, err
}

func (da DataAccess) GetTasksByPages(pool *pgxpool.Pool, queryParams model.QueryParams) (tasks []model.Task, err error) {
	var temp model.Task
	rows, err := pool.Query(context.Background(), "SELECT * FROM tasks WHERE status = $1 LIMIT $2 OFFSET $3", queryParams.Status, queryParams.Limit, queryParams.Offset)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Printf("error while iterating dataset. Error: %s\n", err)
			return tasks, err
		}
		temp.ParseRowsFromTable(values)
		tasks = append(tasks, temp)
	}

	if tasks == nil {
		return tasks, pgx.ErrNoRows
	}
	return tasks, err
}

func (da DataAccess) GetTasksOrderByDate(pool *pgxpool.Pool, queryParams model.QueryParams) (tasks []model.Task, err error) {
	var temp model.Task
	rows, err := pool.Query(context.Background(), "SELECT * FROM tasks WHERE status = $1 ORDER BY date", queryParams.Status)
	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Printf("error while iterating dataset. Error: %s\n", err)
			return tasks, err
		}
		temp.ParseRowsFromTable(values)
		tasks = append(tasks, temp)
	}

	if tasks == nil {
		return tasks, pgx.ErrNoRows
	}
	return tasks, err
}

func (da DataAccess) GetTaskByID(pool *pgxpool.Pool, id int64) (task model.Task, err error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM tasks WHERE id = $1", id)
	if err != nil {
		return model.Task{}, err
	}

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Printf("error while iterating dataset. Error: %s\n", err)
			return task, err
		}
		task.ParseRowsFromTable(values)
	} else {
		return task, pgx.ErrNoRows
	}

	return task, err
}

func (da DataAccess) UpdateTask(pool *pgxpool.Pool, task model.Task) error {
	tag, err := pool.Exec(context.Background(), "UPDATE tasks SET header = $2, description = $3, date = $4, status = $5 WHERE id = $1", task.ID, task.Header, task.Description, task.Date.Time, task.Status)
	if err != nil {
		return err
	}

	if tag.String() == "UPDATE 0" {
		return errors.New("table don't have needed row")
	}

	log.Printf("Task has been updated!%s", tag)

	return err
}

func (da DataAccess) DeleteTask(pool *pgxpool.Pool, id int64) error {
	tag, err := pool.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	if tag.String() == "DELETE 0" {
		return errors.New("table don't have needed row")
	}

	log.Printf("task has been deleted!")

	return err
}
