package service

import (
	"task-go/domain"
	"task-go/dto"
	"task-go/repository"
	"testing"

	"github.com/cockroachdb/errors"
	gomock "github.com/golang/mock/gomock"
)

func Test_projectService_CreateProject(t *testing.T) {

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	testp1, _ := domain.NewProject("test_project1", 1)
	pd1 := &dto.ProjectDto{Name: testp1.GetProjectName(), UserID: uint(testp1.GetUserID())}
	mock1 := repository.NewMockProjectRepositoryIF(mockCtl)
	mock1.EXPECT().CreateProject(pd1).Return(nil)

	testp2, _ := domain.NewProject("test_project2", 2)
	pd2 := &dto.ProjectDto{Name: testp2.GetProjectName(), UserID: uint(testp2.GetUserID())}
	mock2 := repository.NewMockProjectRepositoryIF(mockCtl)
	mock2.EXPECT().CreateProject(pd2).Return(errors.New("test_case_2_error"))

	type fields struct {
		pr repository.ProjectRepositoryIF
	}
	type args struct {
		pr domain.ProjectIF
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock1}, args{testp1}, false},
		{"case2", fields{mock2}, args{testp2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &projectService{
				pr: tt.fields.pr,
			}
			if err := ps.CreateProject(tt.args.pr); (err != nil) != tt.wantErr {
				t.Errorf("projectService.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
