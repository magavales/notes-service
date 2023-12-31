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
	Pool *pgxpool.Pool
}

func NewAccess(pool *pgxpool.Pool) *DataAccess {
	return &DataAccess{Pool: pool}
}

func (da DataAccess) CreateTask(task model.TaskReq) (int64, error) {
	var id int64
	rows, err := da.Pool.Query(context.Background(), "INSERT INTO tasks (header, description, date, status) VALUES ($1, $2, $3, $4) RETURNING id", task.Header, task.Description, task.Date.Time, task.Status)
	if err != nil {
		return 0, err
	}
	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Printf("error while iterating dataset. Error: %s\n", err)
			return 0, err
		}
		id = values[0].(int64)
	} else {
		return 0, pgx.ErrNoRows
	}

	log.Printf("Task has been created!")

	return id, err
}

func (da DataAccess) GetTasks(limit, offset int) (tasks []model.Task, err error) {
	var temp model.Task
	rows, err := da.Pool.Query(context.Background(), "SELECT * FROM tasks LIMIT $1 OFFSET $2", limit, offset)
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

func (da DataAccess) GetTasksWithStatus(status string, limit, offset int) (tasks []model.Task, err error) {
	var temp model.Task
	rows, err := da.Pool.Query(context.Background(), "SELECT * FROM tasks WHERE status = $1 LIMIT $2 OFFSET $3", status, limit, offset)
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

func (da DataAccess) GetTasksOrderBy(status, sort string, limit, offset int) (tasks []model.Task, err error) {
	var temp model.Task
	rows, err := da.Pool.Query(context.Background(), "SELECT * FROM tasks WHERE status = $1 ORDER BY $2 LIMIT $3 OFFSET $4", status, sort, limit, offset)
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

func (da DataAccess) GetTaskByID(id int64) (task model.Task, err error) {
	rows, err := da.Pool.Query(context.Background(), "SELECT * FROM tasks WHERE id = $1", id)
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

func (da DataAccess) UpdateTask(id int64, task model.TaskReq) error {
	tag, err := da.Pool.Exec(context.Background(), "UPDATE tasks SET header = $2, description = $3, date = $4, status = $5 WHERE id = $1", id, task.Header, task.Description, task.Date.Time, task.Status)
	if err != nil {
		return err
	}

	if tag.String() == "UPDATE 0" {
		return errors.New("table don't have needed row")
	}

	log.Printf("Task has been updated!%s", tag)

	return err
}

func (da DataAccess) DeleteTask(id int64) error {
	tag, err := da.Pool.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	if tag.String() == "DELETE 0" {
		return errors.New("table don't have needed row")
	}

	log.Printf("task has been deleted!")

	return err
}
