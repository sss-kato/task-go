package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"task-go/domain"
	"task-go/service"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func Test_projectController_CreateProject(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := service.NewMockProjectServiceIF(mockCtl)

	test_project1, _ := domain.NewProject("test1", 1)
	mock.EXPECT().CreateProject(test_project1).Return(nil)

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/project/create", strings.NewReader(`{"name":"test1","userid":1}`))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	type fields struct {
		ps service.ProjectServiceIF
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &projectController{
				ps: tt.fields.ps,
			}
			if err := pc.CreateProject(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("projectController.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
