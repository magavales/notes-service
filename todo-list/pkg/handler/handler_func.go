package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
)

func (h *Handler) createTask(ctx *gin.Context) {
	var (
		db   database.Database
		task model.Task
		err  error
	)
	err = task.DecodeJSON(ctx.Request.Body)
	if err != nil {
		return
	}

	db.Connect()
	err = db.Access.CreateTask(db.Pool, task)
	if err != nil {
		return
	}
}

func (h *Handler) getTasks(ctx *gin.Context) {
	var (
		db    database.Database
		tasks []model.Task
		err   error
	)

	db.Connect()
	tasks, err = db.Access.GetTasks(db.Pool)
	if err != nil {
		return
	}

	for _, value := range tasks {
		fmt.Println("%s", value)
	}
}

func (h *Handler) getTaskByID(ctx *gin.Context) {
	var (
		db   database.Database
		task model.Task
		err  error
	)
	id, _ := strconv.Atoi(ctx.Param("id"))
	task.ID = int64(id)

	db.Connect()
	task, err = db.Access.GetTaskByID(db.Pool, task.ID)
	if err != nil {
		return
	}
}

func (h *Handler) updateTaskByID(ctx *gin.Context) {
	var (
		db   database.Database
		task model.Task
		err  error
	)
	err = task.DecodeJSON(ctx.Request.Body)
	if err != nil {
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	task.ID = int64(id)

	db.Connect()
	err = db.Access.UpdateTask(db.Pool, task)
	if err != nil {
		return
	}
}

func (h *Handler) deleteTaskByID(ctx *gin.Context) {
	var (
		db   database.Database
		task model.Task
		err  error
	)
	err = task.DecodeJSON(ctx.Request.Body)
	if err != nil {
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	task.ID = int64(id)

	db.Connect()
	err = db.Access.DeleteTask(db.Pool, task.ID)
	if err != nil {
		return
	}
}
