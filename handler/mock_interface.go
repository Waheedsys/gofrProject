// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=mock_interface.go -package=handler
//

// Package handler is a generated GoMock package.
package handler

import (
	entities "gofrProject/entities"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// AddUsers mocks base method.
func (m *MockUserService) AddUsers(user *entities.Users, ctx *gofr.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUsers", user, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUsers indicates an expected call of AddUsers.
func (mr *MockUserServiceMockRecorder) AddUsers(user, ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUsers", reflect.TypeOf((*MockUserService)(nil).AddUsers), user, ctx)
}

// DeleteUsers mocks base method.
func (m *MockUserService) DeleteUsers(name string, ctx *gofr.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsers", name, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsers indicates an expected call of DeleteUsers.
func (mr *MockUserServiceMockRecorder) DeleteUsers(name, ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsers", reflect.TypeOf((*MockUserService)(nil).DeleteUsers), name, ctx)
}

// GetUsers mocks base method.
func (m *MockUserService) GetUsers(ctx *gofr.Context) ([]entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx)
	ret0, _ := ret[0].([]entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserServiceMockRecorder) GetUsers(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserService)(nil).GetUsers), ctx)
}

// GetUsersByName mocks base method.
func (m *MockUserService) GetUsersByName(name string, ctx *gofr.Context) (entities.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByName", name, ctx)
	ret0, _ := ret[0].(entities.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByName indicates an expected call of GetUsersByName.
func (mr *MockUserServiceMockRecorder) GetUsersByName(name, ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByName", reflect.TypeOf((*MockUserService)(nil).GetUsersByName), name, ctx)
}

// UpdateUsers mocks base method.
func (m *MockUserService) UpdateUsers(name string, updateUser *entities.Users, ctx *gofr.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsers", name, updateUser, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUsers indicates an expected call of UpdateUsers.
func (mr *MockUserServiceMockRecorder) UpdateUsers(name, updateUser, ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsers", reflect.TypeOf((*MockUserService)(nil).UpdateUsers), name, updateUser, ctx)
}
