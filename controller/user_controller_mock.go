// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_controller.go

// Package controller is a generated GoMock package.
package controller

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockUserControllerIF is a mock of UserControllerIF interface.
type MockUserControllerIF struct {
	ctrl     *gomock.Controller
	recorder *MockUserControllerIFMockRecorder
}

// MockUserControllerIFMockRecorder is the mock recorder for MockUserControllerIF.
type MockUserControllerIFMockRecorder struct {
	mock *MockUserControllerIF
}

// NewMockUserControllerIF creates a new mock instance.
func NewMockUserControllerIF(ctrl *gomock.Controller) *MockUserControllerIF {
	mock := &MockUserControllerIF{ctrl: ctrl}
	mock.recorder = &MockUserControllerIFMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserControllerIF) EXPECT() *MockUserControllerIFMockRecorder {
	return m.recorder
}

// Signup mocks base method.
func (m *MockUserControllerIF) Signup(e echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signup indicates an expected call of Signup.
func (mr *MockUserControllerIFMockRecorder) Signup(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockUserControllerIF)(nil).Signup), e)
}
