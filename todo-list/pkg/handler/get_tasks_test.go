package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todo-list/pkg/database"
	"todo-list/pkg/model"
)

func TestGetDefaultTasks(t *testing.T) {
	var (
		expectedData []model.Task
		responseData []model.Task
	)
	customTime := new(model.CustomTime)
	customTime1 := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")
	customTime1.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-23 10:32:56")
	expectedData = []model.Task{
		{
			ID:          1,
			Header:      "Погулять в парке Коломенское",
			Description: "сегодня",
			Date:        *customTime,
			Status:      "uncompleted",
		},
		{
			ID:          2,
			Header:      "Сделать домашнее задание",
			Description: "математика, физика",
			Date:        *customTime1,
			Status:      "completed",
		},
	}
	gin.SetMode(gin.TestMode)
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	router := gin.Default()
	router.GET("/api/v1/tasks", handler.getTasks)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	err = json.Unmarshal(resp.Body.Bytes(), &responseData)
	if err != nil {
		log.Fatalf("Can't unmarshal response. Error: %s", err)
		return
	}

	assert.Equal(t, 200, resp.Code, "Test for default get tasks function\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}

func TestGetTasksWithOffsetAndLimit(t *testing.T) {
	var (
		expectedData []model.Task
		responseData []model.Task
	)
	customTime := new(model.CustomTime)
	customTime1 := new(model.CustomTime)
	customTime2 := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-23 10:32:56")
	customTime1.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 21:43:14")
	customTime2.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-25 22:38:20")
	expectedData = []model.Task{
		{
			ID:          2,
			Header:      "Сделать домашнее задание",
			Description: "математика, физика",
			Date:        *customTime,
			Status:      "completed",
		},
		{
			ID:          3,
			Header:      "Погладить кота",
			Description: "сегодня",
			Date:        *customTime1,
			Status:      "uncompleted",
		},
		{
			ID:          4,
			Header:      "Купить чипсы",
			Description: "очень выкусные",
			Date:        *customTime2,
			Status:      "completed",
		},
	}
	gin.SetMode(gin.TestMode)
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	router := gin.Default()
	router.GET("/api/v1/tasks", handler.getTasks)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks?limit=3&offset=1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	err = json.Unmarshal(resp.Body.Bytes(), &responseData)
	if err != nil {
		log.Fatalf("Can't unmarshal response. Error: %s", err)
		return
	}

	assert.Equal(t, 200, resp.Code, "Test for get tasks function with pagination\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}

func TestGetTasksWithStatus(t *testing.T) {
	var (
		expectedData []model.Task
		responseData []model.Task
	)
	customTime := new(model.CustomTime)
	customTime1 := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-23 10:32:56")
	customTime1.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-25 22:38:20")
	expectedData = []model.Task{
		{
			ID:          2,
			Header:      "Сделать домашнее задание",
			Description: "математика, физика",
			Date:        *customTime,
			Status:      "completed",
		},
		{
			ID:          4,
			Header:      "Купить чипсы",
			Description: "очень выкусные",
			Date:        *customTime1,
			Status:      "completed",
		},
	}
	gin.SetMode(gin.TestMode)
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	router := gin.Default()
	router.GET("/api/v1/tasks", handler.getTasks)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks?status=completed", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	err = json.Unmarshal(resp.Body.Bytes(), &responseData)
	if err != nil {
		log.Fatalf("Can't unmarshal response. Error: %s", err)
		return
	}

	assert.Equal(t, 200, resp.Code, "#1 Test for get all tasks function\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}

func TestGetTasksOrderByStatus(t *testing.T) {
	var (
		expectedData []model.Task
		responseData []model.Task
	)
	customTime := new(model.CustomTime)
	customTime1 := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")
	customTime1.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 21:43:14")
	expectedData = []model.Task{
		{
			ID:          1,
			Header:      "Погулять в парке Коломенское",
			Description: "сегодня",
			Date:        *customTime,
			Status:      "uncompleted",
		},
		{
			ID:          3,
			Header:      "Погладить кота",
			Description: "сегодня",
			Date:        *customTime1,
			Status:      "uncompleted",
		},
	}
	gin.SetMode(gin.TestMode)
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	router := gin.Default()
	router.GET("/api/v1/tasks", handler.getTasks)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks?status=uncompleted&sort=date", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	err = json.Unmarshal(resp.Body.Bytes(), &responseData)
	if err != nil {
		log.Fatalf("Can't unmarshal response. Error: %s", err)
		return
	}

	assert.Equal(t, 200, resp.Code, "#1 Test for get all tasks function\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}
