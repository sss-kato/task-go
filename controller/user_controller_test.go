package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"task-go/domain"
	"task-go/service"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_userController_Signup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := service.NewMockUserServiceIF(mockCtl)

	// test case1
	user1, _ := domain.NewUser("test1", "test1", "test1@gmail.com")
	mock.EXPECT().Signup(user1).Return(domain.UserResponse{ID: 1, Name: "test"}, nil)
	e := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"test1","password":"test1","mailadress":"test1@gmail.com"}`))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c := e.NewContext(req1, rec1)

	// test case2
	user2, _ := domain.NewUser("test2", "test2", "test2@gmail.com")
	mock.EXPECT().Signup(user2).Return(domain.UserResponse{}, errors.New("error"))
	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"test2","password":"test2","mailadress":"test2@gmail.com"}`))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)

	// test case3
	e3 := echo.New()
	req3 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"tes3","password":"test3","mailadress":"test3@gmail.com"}`))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e3.NewContext(req3, rec3)

	// test case4
	e4 := echo.New()
	req4 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"test4","password":"tes4","mailadress":"test4@gmail.com"}`))
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e4.NewContext(req4, rec4)

	// test case5
	e5 := echo.New()
	req5 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"test5","password":"test5","mailadress":"test"}`))
	req5.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec5 := httptest.NewRecorder()
	c5 := e5.NewContext(req5, rec5)

	// test case6
	e6 := echo.New()
	req6 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"test6","password":"test6","mailadress":"test6@@gmail.com"}`))
	req6.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec6 := httptest.NewRecorder()
	c6 := e6.NewContext(req6, rec6)

	recMap := map[string]*httptest.ResponseRecorder{"case1": rec1, "case2": rec2, "case3": rec3, "case4": rec4, "case5": rec5, "case6": rec6}

	type fields struct {
		us service.UserServiceIF
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock}, args{c}, false},
		{"case2", fields{mock}, args{c2}, false},
		{"case3", fields{}, args{c3}, false},
		{"case4", fields{}, args{c4}, false},
		{"case5", fields{}, args{c5}, false},
		{"case6", fields{}, args{c6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userController{
				us: tt.fields.us,
			}
			if err := uc.Signup(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userController.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}

			rec := recMap[tt.name]
			expectedHttpStatusCode := 0
			switch tt.name {
			case "case1":
				expectedHttpStatusCode = http.StatusCreated
			case "case2":
				expectedHttpStatusCode = http.StatusInternalServerError
			case "case3", "case4", "case5", "case6":
				expectedHttpStatusCode = http.StatusBadRequest
			}

			if expectedHttpStatusCode != 0 && rec.Code != expectedHttpStatusCode {
				t.Errorf("expected status %d; got %d", expectedHttpStatusCode, rec.Code)

			}

			expectedJSON := ""
			switch tt.name {
			case "case1":
				expectedJSON = `{"id":1,"name":"test"}`
			case "case2":
				expectedJSON = `{"message":"signup failed."}`
			case "case3":
				expectedJSON = `{"message":"user name must be at least five characters long."}`
			case "case4":
				expectedJSON = `{"message":"password must be at least five characters long."}`
			case "case5":
				expectedJSON = `{"message":"mailadress must be at least five characters long."}`
			case "case6":
				expectedJSON = `{"message":"mailadress is invalid."}`
			}

			if expectedJSON != "" {

				assert.JSONEq(t, expectedJSON, rec.Body.String(), fmt.Sprintf("expected JSON: %s; actual JSON: %s", expectedJSON, rec.Body.String()))
			}

		})
	}
}

func Test_userController_Login(t *testing.T) {

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := service.NewMockUserServiceIF(mockCtl)

	// test case1
	user1, _ := domain.NewUser("test1", "test1", "test1@gmail.com")
	mock.EXPECT().Login(user1).Return("test", nil)
	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"name":"test1","password":"test1","mailadress":"test1@gmail.com"}`))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	// test case2
	user2, _ := domain.NewUser("test2", "test2", "test2@gmail.com")
	mock.EXPECT().Login(user2).Return("", errors.New("test"))
	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"name":"test2","password":"test2","mailadress":"test2@gmail.com"}`))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)

	// test case3
	e3 := echo.New()
	req3 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"name":"tes3","password":"test3","mailadress":"test3@gmail.com"}`))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e3.NewContext(req3, rec3)

	// test case4
	e4 := echo.New()
	req4 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"name":"test4","password":"tes4","mailadress":"test4@gmail.com"}`))
	req4.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec4 := httptest.NewRecorder()
	c4 := e4.NewContext(req4, rec4)

	// test case5
	e5 := echo.New()
	req5 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"name":"test5","password":"test5","mailadress":"test"}`))
	req5.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec5 := httptest.NewRecorder()
	c5 := e5.NewContext(req5, rec5)

	// test case6
	e6 := echo.New()
	req6 := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"name":"test6","password":"test6","mailadress":"test6@@gmail.com"}`))
	req6.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec6 := httptest.NewRecorder()
	c6 := e6.NewContext(req6, rec6)

	recMap := map[string]*httptest.ResponseRecorder{"case1": rec1, "case2": rec2, "case3": rec3, "case4": rec4, "case5": rec5, "case6": rec6}

	type fields struct {
		us service.UserServiceIF
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock}, args{c1}, false},
		{"case2", fields{mock}, args{c2}, false},
		{"case3", fields{mock}, args{c3}, false},
		{"case4", fields{mock}, args{c4}, false},
		{"case5", fields{mock}, args{c5}, false},
		{"case6", fields{mock}, args{c6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userController{
				us: tt.fields.us,
			}
			if err := uc.Login(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userController.Login() error = %v, wantErr %v", err, tt.wantErr)
			}

			rec := recMap[tt.name]
			expectedHttpStatusCode := 0
			switch tt.name {
			case "case1":
				expectedHttpStatusCode = http.StatusOK
			case "case2":
				expectedHttpStatusCode = http.StatusInternalServerError

			case "case3", "case4", "case5", "case6":
				expectedHttpStatusCode = http.StatusBadRequest
			}

			if expectedHttpStatusCode != 0 && rec.Code != expectedHttpStatusCode {
				t.Errorf("expected status %d; got %d", expectedHttpStatusCode, rec.Code)

			}

			if tt.name == "case1" && rec.Result().Cookies()[0].Name != "token" {
				t.Errorf("expected cookie name %s; got %s", "token", rec.Result().Cookies()[0].Name)
			}

			if tt.name == "case1" && rec.Result().Cookies()[0].Value != "test" {
				t.Errorf("expected cookie name %s; got %s", "test", rec.Result().Cookies()[0].Value)
			}

			if tt.name == "case1" && rec.Result().Cookies()[0].Expires.After(time.Now().Add(24*time.Hour)) {
				t.Error("cookie time is wrong")
			}

			if tt.name == "case1" && rec.Result().Cookies()[0].Path != "/" {
				t.Errorf("expected cookie name %s; got %s", "/", rec.Result().Cookies()[0].Path)
			}

			tampJSON := `{"message":"%s"}`
			expectedJSON := ""
			errorMsg := ""
			switch tt.name {

			case "case2":
				errorMsg = domain.ErrorMsg09
			case "case3":
				errorMsg = domain.ErrorMsg02
			case "case4":
				errorMsg = domain.ErrorMsg04
			case "case5":
				errorMsg = domain.ErrorMsg06
			case "case6":
				errorMsg = domain.ErrorMsg07
			}

			expectedJSON = fmt.Sprintf(tampJSON, errorMsg)
			if errorMsg != "" {

				assert.JSONEq(t, expectedJSON, rec.Body.String(), fmt.Sprintf("expected JSON: %s; actual JSON: %s", expectedJSON, rec.Body.String()))
			}

		})
	}
}

func Test_userController_Logout(t *testing.T) {

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := service.NewMockUserServiceIF(mockCtl)

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/logout", strings.NewReader(""))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req1.AddCookie(&http.Cookie{
		Name:  "token",
		Value: "test",
	})
	req1.AddCookie(&http.Cookie{
		Name:  "toke",
		Value: "test",
	})
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/logout", strings.NewReader(""))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req2.AddCookie(&http.Cookie{
		Name:  "toke",
		Value: "test",
	})
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)

	type fields struct {
		us service.UserServiceIF
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock}, args{c1}, false},
		{"case2", fields{mock}, args{c2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userController{
				us: tt.fields.us,
			}
			if err := uc.Logout(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userController.Logout() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.name == "case1" {

				cookie := rec1.Result().Cookies()[0]
				if cookie.MaxAge != -1 {
					t.Error("maxage is wrong")
				}
				if !cookie.Expires.Before(time.Now()) {
					t.Error("expire time is wrong")
				}
			}

			if tt.name == "casa2" {
				expectedJSON := `{"message":"no permission."}`
				assert.JSONEq(t, expectedJSON, rec2.Body.String(), fmt.Sprintf("expected JSON: %s; actual JSON: %s", expectedJSON, rec2.Body.String()))
			}

		})
	}
}
