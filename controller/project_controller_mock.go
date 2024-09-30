// Code generated by MockGen. DO NOT EDIT.
// Source: ./project_controller.go

// Package controller is a generated GoMock package.
package controller

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockProjectControllerIF is a mock of ProjectControllerIF interface.
type MockProjectControllerIF struct {
	ctrl     *gomock.Controller
	recorder *MockProjectControllerIFMockRecorder
}

// MockProjectControllerIFMockRecorder is the mock recorder for MockProjectControllerIF.
type MockProjectControllerIFMockRecorder struct {
	mock *MockProjectControllerIF
}

// NewMockProjectControllerIF creates a new mock instance.
func NewMockProjectControllerIF(ctrl *gomock.Controller) *MockProjectControllerIF {
	mock := &MockProjectControllerIF{ctrl: ctrl}
	mock.recorder = &MockProjectControllerIFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectControllerIF) EXPECT() *MockProjectControllerIFMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockProjectControllerIF) CreateProject(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectControllerIFMockRecorder) CreateProject(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectControllerIF)(nil).CreateProject), c)
}