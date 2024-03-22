package controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"task-go/model"
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
	mock.EXPECT().Signup(&model.User{Name: "test", Password: "test"}).Return(model.UserResponse{ID: 1, Name: "test"}, nil)
	mock.EXPECT().Signup(&model.User{Name: "test2", Password: "test2"}).Return(model.UserResponse{}, errors.New("error"))

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"test","password":"test"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(`{"name":"test2","password":"test2"}`))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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
		{"case1", fields{mock}, args{c}, false},
		{"case2", fields{mock}, args{c2}, true},
		{"case3", fields{mock}, args{c2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &userController{
				us: tt.fields.us,
			}
			if err := uc.Signup(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("userController.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
			if rec.Code != http.StatusCreated {
				t.Errorf("expected status %d; got %d", http.StatusOK, rec.Code)
			}
			if tt.name == "case1" {
				expectedJSON := `{"id":1,"name":"test"}`
				// if rec.Body.String() != expectedJSON {
				// 	t.Errorf("expected body %s; got %s", expectedJSON, rec.Body.String())
				// }

				// assert.JSONEq(t, expectedJSON, rec.Body.String(), t.Errorf("expected body %s; got %s", expectedJSON, rec.Body.String()))
				assert.JSONEq(t, expectedJSON, rec.Body.String(), fmt.Sprintf("expected JSON: %s; actual JSON: %s", expectedJSON, rec.Body.String()))
			}
		})
	}
}
