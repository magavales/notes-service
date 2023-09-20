package handler

import (
	"github.com/gin-gonic/gin"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
)

func (h *Handler) CreateTask(ctx *gin.Context) {
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
