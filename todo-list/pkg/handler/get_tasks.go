package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

// @Summary      Get tasks
// @Description  get tasks
// @Tags         get
// @Accept       json
// @Param        queryPagination   	query      	model.Pagination  	true  "Pagination"
// @Param        queryStatus   		query      	model.Status  		true  "Status"
// @Param        querySort   		query      	model.Sort  		true  "Sort"
// @Success      200  				{object}  	[]model.Task
// @Failure      400  				{object}  	response.Response
// @Failure      404  				{object}  	response.Response
// @Failure      500  				{object}  	response.Response
// @Router       /api/v1/tasks 		[get]
func (h *Handler) getTasks(ctx *gin.Context) {
	var (
		tasks           []model.Task
		resp            response.Response
		queryStatus     model.Status
		queryPagination model.Pagination
		querySort       model.Sort
		err             error
	)
	resp.RespWriter = ctx.Writer

	queryPagination.ParseQueryParams(ctx.Request.URL)
	err = queryStatus.ParseQueryParams(ctx.Request.URL)
	if err != nil {
		tasks, err = h.db.Access.GetTasks(queryPagination.Limit, queryPagination.Offset)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				log.Printf("Tasks aren't in database. Error: %s\n", err)
				resp.SetStatusNotFound()
				return
			} else {
				log.Printf("The service couldn't get tasks from database. Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			}
		}

		jdata, err := json.Marshal(tasks)
		if err != nil {
			log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
			resp.SetStatusInternalServerError()
			return
		}
		resp.SetStatusOk()
		resp.SetData(jdata)
		return
	}
	err = querySort.ParseQueryParams(ctx.Request.URL)
	if err != nil {
		tasks, err = h.db.Access.GetTasksWithStatus(queryStatus.Status, queryPagination.Limit, queryPagination.Offset)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				log.Printf("Tasks aren't in database. Error: %s\n", err)
				resp.SetStatusNotFound()
				return
			} else {
				log.Printf("The service couldn't get tasks from database. Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			}
		}

		jdata, err := json.Marshal(tasks)
		if err != nil {
			log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
			resp.SetStatusInternalServerError()
			return
		}
		resp.SetStatusOk()
		resp.SetData(jdata)
		return
	} else {
		tasks, err = h.db.Access.GetTasksOrderBy(queryStatus.Status, querySort.Sort, queryPagination.Limit, queryPagination.Offset)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				log.Printf("Tasks aren't in database. Error: %s\n", err)
				resp.SetStatusNotFound()
				return
			} else {
				log.Printf("The service couldn't get tasks from database. Error: %s\n", err)
				resp.SetStatusInternalServerError()
				return
			}
		}

		jdata, err := json.Marshal(tasks)
		if err != nil {
			log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
			resp.SetStatusInternalServerError()
			return
		}
		resp.SetStatusOk()
		resp.SetData(jdata)
		return
	}
}
