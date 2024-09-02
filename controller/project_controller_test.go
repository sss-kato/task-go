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

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_projectController_CreateProject(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mock1 := service.NewMockProjectServiceIF(mockCtl)
	test_project1, _ := domain.NewProject("test1", 1)
	mock1.EXPECT().CreateProject(test_project1).Return(nil)

	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/project/create", strings.NewReader(`{"name":"test1","userid":1}`))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)

	mock2 := service.NewMockProjectServiceIF(mockCtl)
	test_project2, _ := domain.NewProject("test1", 1)
	mock2.EXPECT().CreateProject(test_project2).Return(errors.New("error"))

	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/project/create", strings.NewReader(`{"name":"test1","userid":1}`))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)

	recMap := map[string]*httptest.ResponseRecorder{"case1": rec1, "case2": rec2}

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
		{"case1", fields{mock1}, args{c1}, false},
		{"case2", fields{mock2}, args{c2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &projectController{
				ps: tt.fields.ps,
			}

			if err := pc.CreateProject(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("projectController.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
			}

			rec := recMap[tt.name]
			expectedHttpStatusCode := 0
			switch tt.name {
			case "case1":
				expectedHttpStatusCode = http.StatusOK
			case "case2":
				expectedHttpStatusCode = http.StatusInternalServerError
			}

			if expectedHttpStatusCode != 0 && rec.Code != expectedHttpStatusCode {
				t.Errorf("expected status %d; got %d", expectedHttpStatusCode, rec.Code)

			}

			tampJSON := `{"message":"%s"}`
			expectedJSON := ""
			errorMsg := ""
			switch tt.name {

			case "case2":
				errorMsg = domain.ErrorMsg13
			}

			expectedJSON = fmt.Sprintf(tampJSON, errorMsg)

			if errorMsg != "" {
				assert.JSONEq(t, expectedJSON, rec.Body.String(), fmt.Sprintf("expected JSON: %s; actual JSON: %s", expectedJSON, rec.Body.String()))
			}
		})
	}
}
