package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

// @Summary      Update task
// @Description  update task
// @Tags         update
// @Accept       json
// @Param        id   	path      	int  true  "Task ID"
// @Param        task   body      	model.TaskReq  true  "new task"
// @Success      200  {object}  	response.Response
// @Failure      400  {object}  	response.Response
// @Failure      404  {object}  	response.Response
// @Failure      500  {object}  	response.Response
// @Router       /api/v1/tasks/:id 	[put]
func (h *Handler) updateTaskByID(ctx *gin.Context) {
	var (
		id          model.TaskID
		task        model.TaskReq
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

	temp, _ := strconv.Atoi(ctx.Param("id"))
	id.ID = int64(temp)

	err = h.db.Access.UpdateTask(id.ID, task)
	if err != nil {
		if errors.As(err, &updateError) {
			log.Printf("Task with id = %d isn't in database. Error: %s\n", id.ID, err)
			resp.SetStatusNotFound()
			return
		} else {
			log.Printf("The service couldn't update data of task from database with id = %d. Error: %s\n", id.ID, err)
			resp.SetStatusInternalServerError()
			return
		}
	}

	resp.SetStatusOk()
}
