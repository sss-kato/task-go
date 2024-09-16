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
	userMock := controller.NewMockUserControllerIF(mockCtl)
	prjMock := controller.NewMockProjectControllerIF(mockCtl)

	// e1 := echo.New()
	// e1.POST("/signup", userMock.Signup)

	// e2 := echo.New()
	// e2.POST("/login", userMock.Login)

	// e3 := echo.New()
	// e3.POST("/logout", userMock.Logout)

	// cmp := map[string]int{"case1": 0, "case2": 1, "case3": 2, "case4": 3}

	e := echo.New()
	e.POST("/signup", userMock.Signup)
	e.POST("/login", userMock.Login)
	e.POST("/logout", userMock.Logout)
	e.POST("/project/create", prjMock.CreateProject)

	type args struct {
		uc controller.UserControllerIF
		pc controller.ProjectControllerIF
	}
	tests := []struct {
		name string
		args args
		want *echo.Echo
	}{
		// TODO: Add test cases.
		{"case1", args{userMock, prjMock}, e},
		// {"case1", args{userMock, prjMock}, e1},
		// {"case2", args{userMock, prjMock}, e2},
		// {"case3", args{userMock, prjMock}, e3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewRouter(tt.args.uc, tt.args.pc)

			routes := got.Routes()

			// path := routes[0].Path
			// i := cmp[tt.name]
			// path := routes[i].Path
			// method := routes[i].Method
			// name := routes[i].Name
			// switch tt.name {
			// case "case1":
			// 	checkPath(path, "/signup", t)
			// 	checkMethod(method, "POST", t)
			// 	checkName(name, "task-go/controller.UserControllerIF.Signup-fm", t)

			// case "case2":
			// 	checkPath(path, "/login", t)
			// 	checkMethod(method, "POST", t)
			// 	checkName(name, "task-go/controller.UserControllerIF.Login-fm", t)
			// case "case3":
			// 	checkPath(path, "/logout", t)
			// 	checkMethod(method, "POST", t)
			// 	checkName(name, "task-go/controller.UserControllerIF.Logout-fm", t)
			// default:
			// }

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

				case "/project/create":
					if route.Method != "POST" {
						t.Errorf("NewRouter().Method = %v, want %v", route.Method, "POST")
					}
					if route.Name != "task-go/controller.ProjectControllerIF.CreateProject-fm" {
						t.Errorf("NewRouter().Name = %v, want %v", route.Name, "task-go/controller.ProjectControllerIF.CreateProject-fm")
					}

				default:
					t.Errorf("NewRouter() = %v", route.Path)
				}

			}

		})
	}

}

func checkPath(actcualPath string, expectedPath string, t *testing.T) {
	if actcualPath != expectedPath {
		t.Errorf("NewRouter().Path = %v, want %v", actcualPath, expectedPath)
	}
}

func checkMethod(actcualMethod string, extpectedMethod string, t *testing.T) {

	if actcualMethod != extpectedMethod {
		t.Errorf("NewRouter().Method = %v, want %v", actcualMethod, extpectedMethod)
	}
}

func checkName(actualName string, expectedName string, t *testing.T) {

	if actualName != expectedName {
		t.Errorf("NewRouter().Name = %v, want %v", actualName, expectedName)
	}
}
