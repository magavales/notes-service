package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

func (h *Handler) deleteTaskByID(ctx *gin.Context) {
	var (
		db          database.Database
		task        model.Task
		resp        response.Response
		err         error
		deleteError error
	)
	resp.RespWriter = ctx.Writer
	deleteError = errors.New("table don't have needed row")

	id, _ := strconv.Atoi(ctx.Param("id"))
	task.ID = int64(id)

	err = db.Connect()
	if err != nil {
		log.Printf("Service can't connect to database: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}

	err = db.Access.DeleteTask(db.Pool, task.ID)
	if err != nil {
		if errors.As(err, &deleteError) {
			log.Printf("Task with id = %d isn't in database. Error: %s\n", task.ID, err)
			resp.SetStatusNotFound()
			return
		} else {
			log.Printf("The service couldn't delete task from database with id = %d. Error: %s\n", task.ID, err)
			resp.SetStatusInternalServerError()
			return
		}
	}

	resp.SetStatusOk()
}
