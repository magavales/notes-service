package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todo-list/pkg/database"
	mock_database "todo-list/pkg/database/mocks"
	"todo-list/pkg/model"
)

func TestHandler_createTask1(t *testing.T) {
	type mockBehaviour func(s *mock_database.MockAccess, task model.TaskReq)

	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")

	testTable := []struct {
		name                string
		inputBody           string
		inputTask           model.TaskReq
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"header": "Погулять в парке Коломенское", "description": "сегодня", "date": "2023-09-27 16:00:00", "status": "uncompleted"}`,
			inputTask: model.TaskReq{
				Header:      "Погулять в парке Коломенское",
				Description: "сегодня",
				Date:        *customTime,
				Status:      "uncompleted",
			},
			mockBehaviour: func(s *mock_database.MockAccess, task model.TaskReq) {
				s.EXPECT().CreateTask(task).Return(int64(1), nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"task_id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			create := mock_database.NewMockAccess(c)
			testCase.mockBehaviour(create, testCase.inputTask)

			access := &database.Database{Access: create}
			handler := NewHandler(access)

			router := gin.Default()
			router.POST("/api/v1/tasks", handler.createTask)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/tasks", bytes.NewBufferString(testCase.inputBody))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, testCase.expectedStatusCode, resp.Code)
			assert.Equal(t, testCase.expectedRequestBody, resp.Body.String())
		})
	}
}
