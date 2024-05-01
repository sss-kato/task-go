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

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_userController_Signup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := service.NewMockUserServiceIF(mockCtl)

	// test case1
	mock.EXPECT().Signup(domain.NewUser("test1", "test1", "test1@gmail.com")).Return(domain.UserResponse{ID: 1, Name: "test"}, nil)
	e := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"<test1","password":"test1","mailadress":"test1@gmail.com"}`))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c := e.NewContext(req1, rec1)

	// test case2
	mock.EXPECT().Signup(domain.NewUser("test2", "test2", "test2@gmail.com")).Return(domain.UserResponse{}, errors.New("error"))
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
			// if rec.Code != http.StatusCreated {
			// 	t.Errorf("expected status %d; got %d", http.StatusOK, rec.Code)
			// }

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
				expectedJSON = `{"message":"username must be at least five characters long."}`
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
