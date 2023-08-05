// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	models "github.com/nanmenkaimak/to-do-list-region/internal/models"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabaseRepo is a mock of DatabaseRepo interface.
type MockDatabaseRepo struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRepoMockRecorder
}

// MockDatabaseRepoMockRecorder is the mock recorder for MockDatabaseRepo.
type MockDatabaseRepoMockRecorder struct {
	mock *MockDatabaseRepo
}

// NewMockDatabaseRepo creates a new mock instance.
func NewMockDatabaseRepo(ctrl *gomock.Controller) *MockDatabaseRepo {
	mock := &MockDatabaseRepo{ctrl: ctrl}
	mock.recorder = &MockDatabaseRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRepo) EXPECT() *MockDatabaseRepoMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockDatabaseRepo) CreateTask(task models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockDatabaseRepoMockRecorder) CreateTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockDatabaseRepo)(nil).CreateTask), task)
}

// DeleteTask mocks base method.
func (m *MockDatabaseRepo) DeleteTask(id primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockDatabaseRepoMockRecorder) DeleteTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockDatabaseRepo)(nil).DeleteTask), id)
}

// GetAllTask mocks base method.
func (m *MockDatabaseRepo) GetAllTask(status bool) ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTask", status)
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTask indicates an expected call of GetAllTask.
func (mr *MockDatabaseRepoMockRecorder) GetAllTask(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTask", reflect.TypeOf((*MockDatabaseRepo)(nil).GetAllTask), status)
}

// UpdateStatus mocks base method.
func (m *MockDatabaseRepo) UpdateStatus(id primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockDatabaseRepoMockRecorder) UpdateStatus(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockDatabaseRepo)(nil).UpdateStatus), id)
}

// UpdateTask mocks base method.
func (m *MockDatabaseRepo) UpdateTask(updatedTask models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", updatedTask)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockDatabaseRepoMockRecorder) UpdateTask(updatedTask interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockDatabaseRepo)(nil).UpdateTask), updatedTask)
}
