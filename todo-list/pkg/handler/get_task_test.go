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

func TestGetTask1(t *testing.T) {
	var (
		expectedData model.Task
		responseData model.Task
	)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")
	expectedData = model.Task{
		ID:          1,
		Header:      "Погулять в парке Коломенское",
		Description: "сегодня",
		Date:        *customTime,
		Status:      "uncompleted",
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
	router.GET("/api/v1/tasks/:id", handler.getTaskByID)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks/1", nil)
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

	assert.Equal(t, 200, resp.Code, "#1 Test for get task function\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}

func TestGetTask2(t *testing.T) {
	var (
		expectedData model.Task
		responseData model.Task
	)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-23 10:32:56")
	expectedData = model.Task{
		ID:          2,
		Header:      "Сделать домашнее задание",
		Description: "математика, физика",
		Date:        *customTime,
		Status:      "completed",
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
	router.GET("/api/v1/tasks/:id", handler.getTaskByID)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks/2", nil)
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

	assert.Equal(t, 200, resp.Code, "#2 Test for get task function\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}

func TestGetTask3(t *testing.T) {
	var (
		expectedData model.Task
		responseData model.Task
	)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 21:43:14")
	expectedData = model.Task{
		ID:          3,
		Header:      "Погладить кота",
		Description: "сегодня",
		Date:        *customTime,
		Status:      "uncompleted",
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
	router.GET("/api/v1/tasks/:id", handler.getTaskByID)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks/3", nil)
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

	assert.Equal(t, 200, resp.Code, "#3 Test for get task function\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}

func TestGetTask4(t *testing.T) {
	var (
		expectedData model.Task
		responseData model.Task
	)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-25 22:38:20")
	expectedData = model.Task{
		ID:          4,
		Header:      "Купить чипсы",
		Description: "очень выкусные",
		Date:        *customTime,
		Status:      "completed",
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
	router.GET("/api/v1/tasks/:id", handler.getTaskByID)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/tasks/4", nil)
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

	assert.Equal(t, 200, resp.Code, "#4 Test for get task function\nStatus code is right.")
	assert.Equal(t, expectedData, responseData, "Expected data equals response date.")
}
