package handler

import (
	"bytes"
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

func TestUpdate1(t *testing.T) {
	gin.SetMode(gin.TestMode)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-28 16:00:00")
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	reqData := model.Task{
		Header:      "Погулять в парке Коломенское",
		Description: "завтра",
		Date:        *customTime,
		Status:      "uncompleted",
	}

	jdata, err := json.Marshal(reqData)
	if err != nil {
		log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
		return
	}

	router := gin.Default()
	router.PUT("/api/v1/tasks/:id", handler.updateTaskByID)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/tasks/1", bytes.NewReader(jdata))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#1 Test for updating is completed!")
}

func TestUpdate2(t *testing.T) {
	gin.SetMode(gin.TestMode)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-23 10:32:56")
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	reqData := model.Task{
		Header:      "Сделать домашнее задание",
		Description: "математика, физика",
		Date:        *customTime,
		Status:      "completed",
	}

	jdata, err := json.Marshal(reqData)
	if err != nil {
		log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
		return
	}

	router := gin.Default()
	router.PUT("/api/v1/tasks/:id", handler.updateTaskByID)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/tasks/2", bytes.NewReader(jdata))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#2 Test for updating is completed!")
}

func TestUpdate3(t *testing.T) {
	gin.SetMode(gin.TestMode)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 21:43:14")
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	reqData := model.Task{
		Header:      "Погладить кота",
		Description: "сегодня",
		Date:        *customTime,
		Status:      "completed",
	}

	jdata, err := json.Marshal(reqData)
	if err != nil {
		log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
		return
	}

	router := gin.Default()
	router.PUT("/api/v1/tasks/:id", handler.updateTaskByID)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/tasks/3", bytes.NewReader(jdata))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#1 Test for updating is completed!")
}

func TestUpdate4(t *testing.T) {
	gin.SetMode(gin.TestMode)
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-26 22:38:20")
	handler := new(Handler)

	handler.Config = database.Config{
		User:     "postgres",
		Password: "1703",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
		Conns:    "10",
	}

	reqData := model.Task{
		Header:      "Купить чипсы",
		Description: "завтра",
		Date:        *customTime,
		Status:      "uncompleted",
	}

	jdata, err := json.Marshal(reqData)
	if err != nil {
		log.Printf("The service couldn't encode data to JSON file. Error: %s\n", err)
		return
	}

	router := gin.Default()
	router.PUT("/api/v1/tasks/:id", handler.updateTaskByID)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/tasks/4", bytes.NewReader(jdata))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#4 Test for updating is completed!")
}
