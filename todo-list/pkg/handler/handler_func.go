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

func (h *Handler) getTasks(ctx *gin.Context) {
	var (
		db          database.Database
		tasks       []model.Task
		resp        response.Response
		queryParams model.QueryParams
		err         error
	)
	resp.RespWriter = ctx.Writer

	err = db.Connect()
	if err != nil {
		log.Printf("Service can't connect to database: %s\n", err)
		resp.SetStatusInternalServerError()
		return
	}

	if !ctx.Request.URL.Query().Has("status") && ctx.Request.URL.Query().Has("limit") {
		log.Printf("Query doesn't have parameter 'status'.\n")
		resp.SetStatusBadRequest()
		return
	}
	if ctx.Request.URL.Query().Has("status") && !ctx.Request.URL.Query().Has("limit") && !ctx.Request.URL.Query().Has("offset") {
		queryParams.ParseQueryParams(ctx.Request.URL)
		tasks, err = db.Access.GetTasksOrderByDate(db.Pool, queryParams)
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
	}
	if ctx.Request.URL.Query().Has("status") && ctx.Request.URL.Query().Has("limit") {
		queryParams.ParseQueryParams(ctx.Request.URL)
		tasks, err = db.Access.GetTasksByPages(db.Pool, queryParams)
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
	} else {
		tasks, err = db.Access.GetTasks(db.Pool)
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
	}
}

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

func (h *Handler) updateTaskByID(ctx *gin.Context) {
	var (
		db          database.Database
		task        model.Task
		resp        response.Response
		err         error
		syntaxError *json.SyntaxError
		updateError error
	)
	resp.RespWriter = ctx.Writer
	updateError = errors.New("table don't have needed row")

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

	err = db.Connect()
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
