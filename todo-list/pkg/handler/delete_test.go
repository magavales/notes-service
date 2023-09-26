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
	"todo-list/pkg/database"
	mock_database "todo-list/pkg/database/mocks"
	"todo-list/pkg/model"
)

func TestHandler_deleteTask1(t *testing.T) {
	type mockBehaviour func(s *mock_database.MockAccess, id model.TaskID)

	temp := new(model.TaskID)
	temp.ID = 1

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
				s.EXPECT().DeleteTask(id.ID).Return(nil)
			},
			expectedStatusCode: 200,
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
			router.DELETE("/api/v1/tasks/:id", handler.deleteTaskByID)

			req := httptest.NewRequest(http.MethodDelete, "/api/v1/tasks/1", bytes.NewBufferString(strconv.FormatInt(testCase.inputID.ID, 10)))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, testCase.expectedStatusCode, resp.Code)
		})
	}
}
