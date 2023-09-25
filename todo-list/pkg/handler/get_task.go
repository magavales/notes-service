package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

func (h *Handler) getTaskByID(ctx *gin.Context) {
	var (
		db   database.Database
		task model.Task
		resp response.Response
		err  error
	)
	resp.RespWriter = ctx.Writer

	id, _ := strconv.Atoi(ctx.Param("id"))
	task.ID = int64(id)

	err = db.Connect()
	if err != nil {
		log.Printf("Service can't connect to database: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}

	task, err = db.Access.GetTaskByID(db.Pool, task.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("Task with id = %d isn't in database. Error: %s\n", id, err)
			resp.SetStatusNotFound()
			return
		} else {
			log.Printf("The service couldn't get task from database with id = %d. Error: %s\n", task.ID, err)
			resp.SetStatusInternalServerError()
			return
		}
	}

	jdata, err := json.Marshal(task)
	if err != nil {
		log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}
	resp.SetStatusOk()
	resp.SetData(jdata)
}
