package repository

import (
	"task-go/db"
	"task-go/dto"
	"testing"

	"gorm.io/gorm"
)

func Test_projectRepository_CreateProject(t *testing.T) {
	test1ProjectDto := &dto.ProjectDto{Name: "testproject1", UserID: 1}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		pd dto.ProjectDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{db.NewDB()}, args{*test1ProjectDto}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &projectRepository{
				db: tt.fields.db,
			}
			if err := pr.CreateProject(tt.args.pd); (err != nil) != tt.wantErr {
				t.Errorf("projectRepository.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
