package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	mock_repository "github.com/nanmenkaimak/to-do-list-region/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteTask(t *testing.T) {
	type mockBehavior func(s *mock_repository.MockDatabaseRepo, taskID primitive.ObjectID)

	testDelete := []struct {
		name               string
		inputID            string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedStatusBody string
	}{
		{
			name:    "OK",
			inputID: "64ce7955d91d3b7a500363e3",
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, taskID primitive.ObjectID) {
				s.EXPECT().DeleteTask(taskID).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:               "wrong task id",
			inputID:            "64ce7955d91d3b7a500363e",
			mockBehavior:       func(s *mock_repository.MockDatabaseRepo, taskID primitive.ObjectID) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedStatusBody: `{"error":"the provided hex string is not a valid ObjectID"}`,
		},
		{
			name:    "no data with given id",
			inputID: "64ce7955d91d3b7a500363e3",
			mockBehavior: func(s *mock_repository.MockDatabaseRepo, taskID primitive.ObjectID) {
				s.EXPECT().DeleteTask(taskID).Return(errors.New("mongo: no documents in result"))
			},
			expectedStatusCode: http.StatusNotFound,
			expectedStatusBody: `{"error":"mongo: no documents in result"}`,
		},
	}

	for _, tt := range testDelete {
		t.Run(tt.name, func(t *testing.T) {
			taskID, _ := primitive.ObjectIDFromHex(tt.inputID)

			c := gomock.NewController(t)
			defer c.Finish()

			newRepo := mock_repository.NewMockDatabaseRepo(c)
			tt.mockBehavior(newRepo, taskID)

			repo := &Repository{newRepo}
			NewHandlers(repo)

			router := gin.Default()
			router.DELETE("/api/todo-list/tasks/:id", Repo.DeleteTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/todo-list/tasks/%s", tt.inputID),
				bytes.NewBufferString(tt.inputID))

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedStatusBody, w.Body.String())
		})
	}
}
