package handler

import (
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

func TestHandler_getTasks1(t *testing.T) {
	type mockBehaviour func(s *mock_database.MockAccess, limit, offset int)

	temp := new(model.Pagination)
	temp.Limit = 2
	temp.Offset = 0
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")
	customTime1 := new(model.CustomTime)
	customTime1.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-25 18:00:00")

	testTable := []struct {
		name                string
		inputPagination     model.Pagination
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:            "OK",
			inputPagination: *temp,
			mockBehaviour: func(s *mock_database.MockAccess, limit, offset int) {
				s.EXPECT().GetTasks(limit, offset).Return([]model.Task{
					{
						ID:          int64(1),
						Header:      "Погулять в парке Коломенское",
						Description: "сегодня",
						Date:        *customTime,
						Status:      "uncompleted",
					},
					{
						ID:          int64(2),
						Header:      "Погладить кота",
						Description: "сейчас",
						Date:        *customTime1,
						Status:      "uncompleted",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"task_id":1,"header":"Погулять в парке Коломенское","description":"сегодня","date":"2023-09-27 16:00:00","status":"uncompleted"},{"task_id":2,"header":"Погладить кота","description":"сейчас","date":"2023-09-25 18:00:00","status":"uncompleted"}]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newMockAccess := mock_database.NewMockAccess(c)
			testCase.mockBehaviour(newMockAccess, testCase.inputPagination.Limit, testCase.inputPagination.Offset)

			access := &database.Database{Access: newMockAccess}
			handler := NewHandler(access)

			router := gin.Default()
			router.GET("/api/v1/tasks", handler.getTasks)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks?limit=2&offset=0", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, testCase.expectedStatusCode, resp.Code)
			assert.Equal(t, testCase.expectedRequestBody, resp.Body.String())
		})
	}
}

func TestHandler_getTasks2(t *testing.T) {
	type mockBehaviour func(s *mock_database.MockAccess, status string, limit, offset int)

	temp := new(model.Pagination)
	temp.Limit = 2
	temp.Offset = 0
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")
	customTime1 := new(model.CustomTime)
	customTime1.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-25 18:00:00")
	status := new(model.Status)
	status.Status = "uncompleted"

	testTable := []struct {
		name                string
		inputPagination     model.Pagination
		inputStatus         model.Status
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:            "OK",
			inputPagination: *temp,
			inputStatus:     *status,
			mockBehaviour: func(s *mock_database.MockAccess, status string, limit, offset int) {
				s.EXPECT().GetTasksWithStatus(status, limit, offset).Return([]model.Task{
					{
						ID:          int64(1),
						Header:      "Погулять в парке Коломенское",
						Description: "сегодня",
						Date:        *customTime,
						Status:      "uncompleted",
					},
					{
						ID:          int64(2),
						Header:      "Погладить кота",
						Description: "сейчас",
						Date:        *customTime1,
						Status:      "uncompleted",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"task_id":1,"header":"Погулять в парке Коломенское","description":"сегодня","date":"2023-09-27 16:00:00","status":"uncompleted"},{"task_id":2,"header":"Погладить кота","description":"сейчас","date":"2023-09-25 18:00:00","status":"uncompleted"}]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newMockAccess := mock_database.NewMockAccess(c)
			testCase.mockBehaviour(newMockAccess, testCase.inputStatus.Status, testCase.inputPagination.Limit, testCase.inputPagination.Offset)

			access := &database.Database{Access: newMockAccess}
			handler := NewHandler(access)

			router := gin.Default()
			router.GET("/api/v1/tasks", handler.getTasks)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks?status=uncompleted&limit=2&offset=0", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, testCase.expectedStatusCode, resp.Code)
			assert.Equal(t, testCase.expectedRequestBody, resp.Body.String())
		})
	}
}

func TestHandler_getTasks3(t *testing.T) {
	type mockBehaviour func(s *mock_database.MockAccess, status, sort string, limit, offset int)

	temp := new(model.Pagination)
	temp.Limit = 2
	temp.Offset = 0
	customTime := new(model.CustomTime)
	customTime.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-27 16:00:00")
	customTime1 := new(model.CustomTime)
	customTime1.Time, _ = time.Parse("2006-01-02 15:04:05", "2023-09-25 18:00:00")
	status := new(model.Status)
	status.Status = "uncompleted"
	sort := new(model.Sort)
	sort.Sort = "date"

	testTable := []struct {
		name                string
		inputPagination     model.Pagination
		inputStatus         model.Status
		inputSort           model.Sort
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:            "OK",
			inputPagination: *temp,
			inputStatus:     *status,
			inputSort:       *sort,
			mockBehaviour: func(s *mock_database.MockAccess, status, sort string, limit, offset int) {
				s.EXPECT().GetTasksOrderBy(status, sort, limit, offset).Return([]model.Task{
					{
						ID:          int64(2),
						Header:      "Погладить кота",
						Description: "сейчас",
						Date:        *customTime1,
						Status:      "uncompleted",
					},
					{
						ID:          int64(1),
						Header:      "Погулять в парке Коломенское",
						Description: "сегодня",
						Date:        *customTime,
						Status:      "uncompleted",
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"task_id":2,"header":"Погладить кота","description":"сейчас","date":"2023-09-25 18:00:00","status":"uncompleted"},{"task_id":1,"header":"Погулять в парке Коломенское","description":"сегодня","date":"2023-09-27 16:00:00","status":"uncompleted"}]`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newMockAccess := mock_database.NewMockAccess(c)
			testCase.mockBehaviour(newMockAccess, testCase.inputStatus.Status, testCase.inputSort.Sort, testCase.inputPagination.Limit, testCase.inputPagination.Offset)

			access := &database.Database{Access: newMockAccess}
			handler := NewHandler(access)

			router := gin.Default()
			router.GET("/api/v1/tasks", handler.getTasks)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/tasks?status=uncompleted&sort=date&limit=2&offset=0", nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, testCase.expectedStatusCode, resp.Code)
			assert.Equal(t, testCase.expectedRequestBody, resp.Body.String())
		})
	}
}
