package service

import (
	"task-go/domain"
	"task-go/repository"
	"testing"
)

func Test_projectService_CreateProject(t *testing.T) {
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
