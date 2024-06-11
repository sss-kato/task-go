package router

import (
	"task-go/controller"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestNewRouter(t *testing.T) {

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := controller.NewMockUserControllerIF(mockCtl)

	e := echo.New()
	e.POST("/signup", mock.Signup)
	e.POST("/login", mock.Login)
	e.POST("/logout", mock.Logout)

	type args struct {
		uc controller.UserControllerIF
	}
	tests := []struct {
		name string
		args args
		want *echo.Echo
	}{
		// TODO: Add test cases.
		{"case1", args{mock}, e},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewRouter(tt.args.uc)

			routes := got.Routes()

			for _, route := range routes {
				switch route.Path {

				case "/signup":
					if route.Method != "POST" {
						t.Errorf("NewRouter().Method = %v, want %v", route.Method, "POST")
					}

					if route.Name != "task-go/controller.UserControllerIF.Signup-fm" {
						t.Errorf("NewRouter().Name = %v, want %v", route.Name, "task-go/controller.UserControllerIF.Signup-fm")
					}

				case "/login":
					if route.Method != "POST" {
						t.Errorf("NewRouter().Method = %v, want %v", route.Method, "POST")
					}
					if route.Name != "task-go/controller.UserControllerIF.Login-fm" {
						t.Errorf("NewRouter().Name = %v, want %v", route.Name, "task-go/controller.UserControllerIF.Login-fm")
					}

				case "/logout":
					if route.Method != "POST" {
						t.Errorf("NewRouter().Method = %v, want %v", route.Method, "POST")
					}
					if route.Name != "task-go/controller.UserControllerIF.Logout-fm" {
						t.Errorf("NewRouter().Name = %v, want %v", route.Name, "task-go/controller.UserControllerIF.Login-fm")
					}

				default:
					t.Errorf("NewRouter() = %v", route.Path)
				}

			}

		})
	}
}
