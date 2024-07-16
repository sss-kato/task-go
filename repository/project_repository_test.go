package repository

import (
	"regexp"
	"task-go/dto"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cockroachdb/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_projectRepository_CreateProject(t *testing.T) {

	// gormDB := createDbMockCon(t)
	dbm, mock1, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbm.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 dbm,
		PreferSimpleProtocol: true,
	})
	gormDB1, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing a gorm db connection", err)
	}

	// test case1
	test1ProjectDto := &dto.ProjectDto{Name: "testproject1", UserID: 1}
	// mock1 := createDbMock(t)
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

	// test case 2
	// dbm2, mock2, err := sqlmock.New()
	dbm2, mock2, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbm.Close()
	dialector2 := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 dbm2,
		PreferSimpleProtocol: true,
	})
	gormDB2, err := gorm.Open(dialector2, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing a gorm db connection", err)
	}

	test2ProjectDto := &dto.ProjectDto{Name: "testproject2", UserID: 2}

	// mock2 := createDbMock(t)
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
		pd dto.ProjectDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{gormDB1}, args{*test1ProjectDto}, false},
		{"case2", fields{gormDB2}, args{*test2ProjectDto}, true},
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

func createDbMockCon(t *testing.T) *gorm.DB {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing a gorm db connection", err)
	}

	return gormDB
}

func createDbMock(t *testing.T) sqlmock.Sqlmock {

	_, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return mock

}
