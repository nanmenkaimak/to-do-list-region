package handlers

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	mock_repository "github.com/nanmenkaimak/to-do-list-region/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasks(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockDatabaseRepo, status bool)

	testGet := []struct {
		name               string
		inputUser          string
		status             bool
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedStatusBody string
	}{
		{
			name:      "OK",
			inputUser: "active",
			status:    false,
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, status bool) {
				s.EXPECT().GetAllTasks(status).Return([]models.Task{
					{
						Title:    "something",
						ActiveAt: "2023-08-06",
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedStatusBody: `[{"title":"ВЫХОДНОЙ - something","activeAt":"2023-08-06"}]`,
		},
		{
			name:      "empty result",
			inputUser: "active",
			status:    false,
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, status bool) {
				s.EXPECT().GetAllTasks(status).Return(nil, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedStatusBody: `[]`,
		},
		{
			name:   "some error",
			status: false,
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, status bool) {
				s.EXPECT().GetAllTasks(status).Return(nil, errors.New("some error"))
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"some error"}`,
		},
		{
			name:      "parsing date",
			inputUser: "done",
			status:    true,
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, status bool) {
				s.EXPECT().GetAllTasks(status).Return([]models.Task{
					{
						Title:    "something",
						ActiveAt: "not a time",
					},
				}, nil)
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"parsing time \"not a time\" as \"2006-01-02\": cannot parse \"not a time\" as \"2006\""}`,
		},
	}

	for _, tt := range testGet {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newRepo := mock_repository.NewMockDatabaseRepo(c)
			tt.mockBehavior(newRepo, tt.status)

			repo := &Repository{newRepo}
			NewHandlers(repo)

			router := gin.Default()
			router.GET("/api/todo-list/tasks", Repo.GetTasks)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/todo-list/tasks",
				bytes.NewBufferString(tt.inputUser))

			q := req.URL.Query()
			q.Add("status", tt.inputUser)
			req.URL.RawQuery = q.Encode()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedStatusBody, w.Body.String())
		})
	}
}
