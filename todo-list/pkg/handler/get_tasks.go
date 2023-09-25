package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"log"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
	"todo-list/pkg/response"
)

func (h *Handler) getTasks(ctx *gin.Context) {
	var (
		db              database.Database
		tasks           []model.Task
		resp            response.Response
		queryStatus     model.Status
		queryPagination model.Pagination
		querySort       model.Sort
		err             error
	)
	resp.RespWriter = ctx.Writer

	err = db.Connect(h.Config)
	if err != nil {
		log.Printf("Service can't connect to database: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}
	queryPagination.ParseQueryParams(ctx.Request.URL)
	err = queryStatus.ParseQueryParams(ctx.Request.URL)
	if err != nil {
		tasks, err = db.Access.GetTasks(db.Pool, queryPagination.Limit, queryPagination.Offset)
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
		tasks, err = db.Access.GetTasksWithStatus(db.Pool, queryStatus.Status, queryPagination.Limit, queryPagination.Offset)
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
		tasks, err = db.Access.GetTasksOrderBy(db.Pool, queryStatus.Status, querySort.Sort, queryPagination.Limit, queryPagination.Offset)
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
