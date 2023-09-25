package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

func (h *Handler) createTask(ctx *gin.Context) {
	var (
		db          database.Database
		task        model.Task
		resp        response.Response
		err         error
		syntaxError *json.SyntaxError
	)
	resp.RespWriter = ctx.Writer

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

	err = db.Connect()
	if err != nil {
		log.Printf("Service can't connect to database: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}

	err = db.Access.CreateTask(db.Pool, &task)
	if err != nil {
		log.Printf("The service couldn't create task in database. Error: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	} else {
		respID := new(model.IDofCreatedTask)
		respID.ID = task.ID
		jdata, err := json.Marshal(respID)
		if err != nil {
			log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
			resp.SetStatusInternalServerError()
			return
		}
		resp.SetStatusOk()
		resp.SetData(jdata)
	}
}
