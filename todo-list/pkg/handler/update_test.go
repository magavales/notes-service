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

func TestHandler_updateTask1(t *testing.T) {
	type mockBehaviour func(s *mock_database.MockAccess, id model.TaskID, task model.TaskReq)

	temp := new(model.TaskID)
	temp.ID = 1
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")

	testTable := []struct {
		name                string
		inputID             model.TaskID
		inputBody           string
		inputTask           model.TaskReq
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputID:   *temp,
			inputBody: `{"header": "Погулять в парке Коломенское", "description": "сегодня", "date": "2023-09-27 16:00:00", "status": "uncompleted"}`,
			inputTask: model.TaskReq{
				Header:      "Погулять в парке Коломенское",
				Description: "сегодня",
				Date:        *customTime,
				Status:      "uncompleted",
			},
			mockBehaviour: func(s *mock_database.MockAccess, id model.TaskID, task model.TaskReq) {
				s.EXPECT().UpdateTask(id.ID, task).Return(nil)
			},
			expectedStatusCode: 200,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			update := mock_database.NewMockAccess(c)
			testCase.mockBehaviour(update, testCase.inputID, testCase.inputTask)

			access := &database.Database{Access: update}
			handler := NewHandler(access)

			router := gin.Default()
			router.PUT("/api/v1/tasks/:id", handler.updateTaskByID)

			req := httptest.NewRequest(http.MethodPut, "/api/v1/tasks/1", bytes.NewBufferString(testCase.inputBody))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, testCase.expectedStatusCode, resp.Code)
		})
	}
}
