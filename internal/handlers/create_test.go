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

func TestCreateTask(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockDatabaseRepo, newTask models.Task)

	testCreate := []struct {
		name               string
		inputBody          string
		inputUser          models.Task
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedStatusBody string
	}{
		{
			name:      "OK",
			inputBody: `{"title":"Купить книгу","activeAt":"2023-08-12"}`,
			inputUser: models.Task{
				Title:    "Купить книгу",
				ActiveAt: "2023-08-12",
			},
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {
				s.EXPECT().CreateTask(newTask).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:               "no input json",
			inputBody:          ``,
			mockBehavior:       func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"EOF"}`,
		},
		{
			name:      "missing fields",
			inputBody: `{"activeAt":"2023-08-12"}`,
			inputUser: models.Task{
				ActiveAt: "2023-08-12",
			},
			mockBehavior:       func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"missing values"}`,
		},
		{
			name:      "wrong date",
			inputBody: `{"title":"Купить книгу", "activeAt":"2023-08-40"}`,
			inputUser: models.Task{
				ActiveAt: "2023-08-40",
			},
			mockBehavior:       func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"parsing time \"2023-08-40\": day out of range"}`,
		},
		{
			name:      "similar value",
			inputBody: `{"title":"Купить книгу","activeAt":"2023-08-13"}`,
			inputUser: models.Task{
				Title:    "Купить книгу",
				ActiveAt: "2023-08-13",
			},
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {
				s.EXPECT().CreateTask(newTask).Return(errors.New("there are already similar task"))
			},
			expectedStatusCode: http.StatusNotFound,
			expectedStatusBody: `{"error":"there are already similar task"}`,
		},
	}

	for _, tt := range testCreate {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newRepo := mock_repository.NewMockDatabaseRepo(c)
			tt.mockBehavior(newRepo, tt.inputUser)

			repo := &Repository{newRepo}
			NewHandlers(repo)

			router := gin.Default()
			router.POST("/api/todo-list/tasks", Repo.CreateTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/todo-list/tasks",
				bytes.NewBufferString(tt.inputBody))

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedStatusBody, w.Body.String())
		})
	}
}
