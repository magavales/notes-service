package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

func (h *Handler) updateTaskByID(ctx *gin.Context) {
	var (
		db          database.Database
		task        model.Task
		resp        response.Response
		err         error
		syntaxError *json.SyntaxError
	)
	resp.RespWriter = ctx.Writer
	var updateError any = errors.New("table don't have needed row")

	err = task.DecodeJSON(ctx.Request.Body)
	if err != nil {
		if errors.As(err, &syntaxError) {
			log.Printf("JSON file has syntax error. Error: %s\n", err)
			resp.SetStatusBadRequest()
			return
		} else {
			log.Printf("The service couldn't decode JSON file. Error: %s\n", err)
			resp.SetStatusInternalServerError()
			return
		}
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	task.ID = int64(id)

	err = db.Connect(h.Config)
	if err != nil {
		log.Printf("Service can't connect to database: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}

	err = db.Access.UpdateTask(db.Pool, task)
	if err != nil {
		if errors.As(err, &updateError) {
			log.Printf("Task with id = %d isn't in database. Error: %s\n", task.ID, err)
			resp.SetStatusNotFound()
			return
		} else {
			log.Printf("The service couldn't update data of task from database with id = %d. Error: %s\n", task.ID, err)
			resp.SetStatusInternalServerError()
			return
		}
	}

	resp.SetStatusOk()
}
