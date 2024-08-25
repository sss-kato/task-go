package controller

import (
	"task-go/service"
	"testing"

	"github.com/labstack/echo/v4"
)

func Test_projectController_CreateProject(t *testing.T) {
	// mockCtl := gomock.NewController(t)
	// defer mockCtl.Finish()
	// mock := service.NewMockProjectServiceIF(mockCtl)

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
