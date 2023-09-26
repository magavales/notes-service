package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"todo-list/pkg/database"
	mock_database "todo-list/pkg/database/mocks"
	"todo-list/pkg/model"
)

func TestHandler_getTask1(t *testing.T) {
	type mockBehaviour func(s *mock_database.MockAccess, id model.TaskID)

	temp := new(model.TaskID)
	temp.ID = 1
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")

	testTable := []struct {
		name                string
		inputID             model.TaskID
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:    "OK",
			inputID: *temp,
			mockBehaviour: func(s *mock_database.MockAccess, id model.TaskID) {
				s.EXPECT().GetTaskByID(id.ID).Return(model.Task{
					ID:          int64(1),
					Header:      "Погулять в парке Коломенское",
					Description: "сегодня",
					Date:        *customTime,
					Status:      "uncompleted",
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"task_id":1,"header":"Погулять в парке Коломенское","description":"сегодня","date":"2023-09-27 16:00:00","status":"uncompleted"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newMockAccess := mock_database.NewMockAccess(c)
			testCase.mockBehaviour(newMockAccess, testCase.inputID)

			access := &database.Database{Access: newMockAccess}
			handler := NewHandler(access)

			router := gin.Default()
			router.GET("/api/v1/tasks/:id", handler.getTaskByID)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks/1", bytes.NewBufferString(strconv.FormatInt(testCase.inputID.ID, 10)))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, testCase.expectedStatusCode, resp.Code)
			assert.Equal(t, testCase.expectedRequestBody, resp.Body.String())
		})
	}
}
