package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/pkg/database"
)

func TestDelete1(t *testing.T) {
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
	router.DELETE("/api/v1/tasks/:id", handler.deleteTaskByID)

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/tasks/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#1 Test for updating is completed!")
}

func TestDelete2(t *testing.T) {
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
	router.DELETE("/api/v1/tasks/:id", handler.deleteTaskByID)

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/tasks/2", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#2 Test for updating is completed!")
}

func TestDelete3(t *testing.T) {
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
	router.DELETE("/api/v1/tasks/:id", handler.deleteTaskByID)

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/tasks/3", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#1 Test for updating is completed!")
}

func TestDelete4(t *testing.T) {
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
	router.DELETE("/api/v1/tasks/:id", handler.deleteTaskByID)

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/tasks/4", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code, "#4 Test for updating is completed!")
}
