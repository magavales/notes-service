package database

import (
	"todo-list/pkg/database/table"
	"todo-list/pkg/model"
)

//go:generate mockgen -source=database.go -destination=mocks/mock.go

type Access interface {
	CreateTask(task model.TaskReq) (int64, error)
	GetTasks(limit, offset int) (tasks []model.Task, err error)
	GetTasksWithStatus(status string, limit, offset int) (tasks []model.Task, err error)
	GetTasksOrderBy(status, sort string, limit, offset int) (tasks []model.Task, err error)
	GetTaskByID(id int64) (task model.Task, err error)
	UpdateTask(id int64, task model.TaskReq) error
	DeleteTask(id int64) error
}

type Database struct {
	Access
}

func NewConn(access *table.DataAccess) *Database {
	return &Database{
		Access: table.NewAccess(access.Pool),
	}
}

/*func (db *Database) Connect(conf Config) (*pgxpool.Pool, error) {
	poolConn, err := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s pool_max_conns=%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Name, conf.Conns))
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConn)
	if err != nil {
		return nil, err
	}
	return pool, err
}*/
