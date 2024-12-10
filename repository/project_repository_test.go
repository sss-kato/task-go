package repository

import (
	"regexp"
	"task-go/db"
	"task-go/dto"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

func Test_projectRepository_CreateProject(t *testing.T) {

	// test case1
	mock1, gormDB1 := db.CreateDBMock()
	test1ProjectDto := &dto.ProjectDto{Name: "testproject1", UserID: 1}
	mock1.ExpectBegin()
	resp1 := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock1.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "projects" ("name","user_id","created_at","updated_at") VALUES ($1,$2,$3,$4)`)).
		WithArgs(test1ProjectDto.Name, test1ProjectDto.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(resp1)

	respm1 := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock1.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "project_user_mappings" ("project_id","user_id") VALUES ($1,$2)`)).
		WithArgs(1, test1ProjectDto.UserID).
		WillReturnRows(respm1)

	mock1.ExpectCommit()

	// test case2
	mock2, gormDB2 := db.CreateDBMock()
	test2ProjectDto := &dto.ProjectDto{Name: "testproject2", UserID: 2}
	mock2.ExpectBegin()
	rows2 := sqlmock.NewRows([]string{"id"}).AddRow(5)
	mock2.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "projects" ("name","user_id","created_at","updated_at") VALUES ($1,$2,$3,$4)`)).
		WithArgs(test2ProjectDto.Name, test2ProjectDto.UserID, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows2)
	mock2.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "project_user_mappings" ("project_id","user_id") VALUES ($1,$2)`)).
		WithArgs(5, test2ProjectDto.UserID).
		WillReturnError(errors.New("insertion error"))
	mock2.ExpectRollback()

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		pd *dto.ProjectDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{gormDB1}, args{test1ProjectDto}, false},
		{"case2", fields{gormDB2}, args{test2ProjectDto}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := &projectRepository{
				db: tt.fields.db,
			}
			if err := pr.CreateProject(tt.args.pd); (err != nil) != tt.wantErr {
				t.Errorf("projectRepository.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.name == "case2" {
				if err := mock2.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
			}
		})
	}
}

// func Test_projectRepository_GetProjects(t *testing.T) {

// 	userTest1 := dto.UserDto{ID: 1}

// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		ud dto.UserDto
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    []dto.ProjectDto
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"case1", fields{db.NewDB()}, args{userTest1}, nil, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			pr := &projectRepository{
// 				db: tt.fields.db,
// 			}
// 			got, err := pr.GetProjects(tt.args.ud)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("projectRepository.GetProjects() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("projectRepository.GetProjects() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
