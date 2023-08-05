package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/to-do-list-region/internal/models"
	mock_repository "github.com/nanmenkaimak/to-do-list-region/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateTask(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockDatabaseRepo, newTask models.Task)

	testUpdate := []struct {
		name               string
		inputBody          string
		inputUser          models.Task
		inputID            string
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
			inputID: "64ce7955d91d3b7a500363e3",
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {
				s.EXPECT().UpdateTask(newTask).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:               "no input json",
			inputBody:          ``,
			inputID:            "64ce7955d91d3b7a500363e3",
			mockBehavior:       func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"EOF"}`,
		},
		{
			name:      "wrong task id",
			inputBody: `{"title":"Купить книгу","activeAt":"2023-08-12"}`,
			inputUser: models.Task{
				Title:    "Купить книгу",
				ActiveAt: "2023-08-12",
			},
			inputID:            "64ce7955d91d3b7a500363e",
			mockBehavior:       func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"the provided hex string is not a valid ObjectID"}`,
		},
		{
			name:      "similar value",
			inputBody: `{"title":"Купить книгу","activeAt":"2023-08-12"}`,
			inputUser: models.Task{
				Title:    "Купить книгу",
				ActiveAt: "2023-08-12",
			},
			inputID: "64ce7955d91d3b7a500363e3",
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, newTask models.Task) {
				s.EXPECT().UpdateTask(newTask).Return(errors.New("there are already similar task"))
			},
			expectedStatusCode: http.StatusNotFound,
			expectedStatusBody: `{"error":"there are already similar task"}`,
		},
	}
	for _, tt := range testUpdate {
		t.Run(tt.name, func(t *testing.T) {
			taskID, _ := primitive.ObjectIDFromHex(tt.inputID)
			tt.inputUser.ID = taskID

			c := gomock.NewController(t)
			defer c.Finish()

			newRepo := mock_repository.NewMockDatabaseRepo(c)
			tt.mockBehavior(newRepo, tt.inputUser)

			repo := &Repository{newRepo}
			NewHandlers(repo)

			router := gin.Default()
			router.PUT("/api/todo-list/tasks/:id", Repo.UpdateTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/todo-list/tasks/%s", tt.inputID),
				bytes.NewBufferString(tt.inputBody))

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedStatusBody, w.Body.String())
		})
	}
}
