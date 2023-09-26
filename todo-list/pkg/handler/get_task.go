package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

// @Summary      Get task
// @Description  get task
// @Tags         get
// @Accept       json
// @Produce		 json
// @Param        id   path      	int  true  "Task ID"
// @Success      200  {object}  	model.Task
// @Failure      400  {object}  	response.Response
// @Failure      404  {object}  	response.Response
// @Failure      500  {object}  	response.Response
// @Router       /api/v1/tasks/:id 	[get]
func (h *Handler) getTaskByID(ctx *gin.Context) {
	var (
		task model.Task
		id   model.TaskID
		resp response.Response
		err  error
	)
	resp.RespWriter = ctx.Writer

	temp, _ := strconv.Atoi(ctx.Param("id"))
	id.ID = int64(temp)

	task, err = h.db.Access.GetTaskByID(id.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("Task with id = %d isn't in database. Error: %s\n", id, err)
			resp.SetStatusNotFound()
			return
		} else {
			log.Printf("The service couldn't get task from database with id = %d. Error: %s\n", id, err)
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
