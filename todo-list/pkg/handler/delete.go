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

// @Summary      Delete task
// @Description  delete task
// @Tags         delete
// @Accept       json
// @Param        id   path      	int  true  "Task ID"
// @Success      200  {object}  	response.Response
// @Failure      400  {object}  	response.Response
// @Failure      404  {object}  	response.Response
// @Failure      500  {object}  	response.Response
// @Router       /api/v1/tasks/:id 	[delete]
func (h *Handler) deleteTaskByID(ctx *gin.Context) {
	var (
		db   database.Database
		id   model.TaskID
		task model.Task
		resp response.Response
		err  error
	)
	resp.RespWriter = ctx.Writer
	var deleteError any = errors.New("table don't have needed row")

	temp, _ := strconv.Atoi(ctx.Param("id"))
	id.ID = int64(temp)

	err = db.Connect(h.Config)
	if err != nil {
		log.Printf("Service can't connect to database: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}

	err = db.Access.DeleteTask(db.Pool, id.ID)
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
