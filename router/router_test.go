package router

// func TestNewRouter(t *testing.T) {

// 	mockCtl := gomock.NewController(t)
// 	defer mockCtl.Finish()
// 	userMock := controller.NewMockUserControllerIF(mockCtl)

// 	prjMock := controller.NewMockProjectControllerIF(mockCtl)

// 	e := echo.New()
// 	e.POST("/signup", userMock.Signup)
// 	e.POST("/login", userMock.Login)
// 	e.POST("/logout", userMock.Logout)

// 	type args struct {
// 		uc controller.UserControllerIF
// 		pc controller.ProjectControllerIF
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *echo.Echo
// 	}{
// 		// TODO: Add test cases.
// 		{"case1", args{userMock, prjMock}, e},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			got := NewRouter(tt.args.uc, tt.args.pc)

// 			routes := got.Routes()

// 			path := routes[0].Path
// 			method := routes[0].Method
// 			name := routes[0].Name
// 			switch tt.name {
// 			case "case1":
// 				checkPath(path, "/signnup", t)

// 				checkMethod(method, "POST", t)

// 				checkName(name, "task-go/controller.UserControllerIF.Signup-fm", t)

// 			case "case2":

// 			default:
// 			}

// 			for _, route := range routes {
// 				switch route.Path {

// 				case "/signup":
// 					if route.Method != "POST" {
// 						t.Errorf("NewRouter().Method = %v, want %v", route.Method, "POST")
// 					}

// 					if route.Name != "task-go/controller.UserControllerIF.Signup-fm" {
// 						t.Errorf("NewRouter().Name = %v, want %v", route.Name, "task-go/controller.UserControllerIF.Signup-fm")
// 					}

// 				case "/login":
// 					if route.Method != "POST" {
// 						t.Errorf("NewRouter().Method = %v, want %v", route.Method, "POST")
// 					}
// 					if route.Name != "task-go/controller.UserControllerIF.Login-fm" {
// 						t.Errorf("NewRouter().Name = %v, want %v", route.Name, "task-go/controller.UserControllerIF.Login-fm")
// 					}

// 				case "/logout":
// 					if route.Method != "POST" {
// 						t.Errorf("NewRouter().Method = %v, want %v", route.Method, "POST")
// 					}
// 					if route.Name != "task-go/controller.UserControllerIF.Logout-fm" {
// 						t.Errorf("NewRouter().Name = %v, want %v", route.Name, "task-go/controller.UserControllerIF.Login-fm")
// 					}

// 				default:
// 					t.Errorf("NewRouter() = %v", route.Path)
// 				}

// 			}

// 		})
// 	}

// }

// func checkPath(actcualPath string, expectedPath string, t *testing.T) {
// 	if actcualPath != expectedPath {
// 		t.Errorf("NewRouter().Path = %v, want %v", actcualPath, expectedPath)
// 	}
// }

// func checkMethod(actcualMethod string, extpectedMethod string, t *testing.T) {

// 	if actcualMethod != extpectedMethod {
// 		t.Errorf("NewRouter().Method = %v, want %v", actcualMethod, extpectedMethod)
// 	}
// }

// func checkName(actualName string, expectedName string, t *testing.T) {

// 	if actualName != expectedName {
// 		t.Errorf("NewRouter().Name = %v, want %v", actualName, expectedName)
// 	}
// }
