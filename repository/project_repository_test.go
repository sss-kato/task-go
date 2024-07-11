package repository

// func Test_projectRepository_CreateProject(t *testing.T) {
// 	jst, _ := time.LoadLocation("Asia/Tokyo")
// 	now := time.Now().In(jst)
// 	// formattedNow := now.Format("2006-01-02 15:04:05.000")
// 	test1ProjectDto := &dto.ProjectDto{Name: "testproject1", UserID: 1}
// 	test2ProjectDto := &dto.ProjectDto{Name: "testproject2", UserID: 2, CreatedAt: now, UpdatedAt: now}

// 	dbm, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer dbm.Close()
// 	dialector := postgres.New(postgres.Config{
// 		DSN:                  "sqlmock_db_0",
// 		DriverName:           "postgres",
// 		Conn:                 dbm,
// 		PreferSimpleProtocol: true,
// 	})
// 	gormDB, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when initializing a gorm db connection", err)
// 	}

// 	mock.ExpectBegin()

// 	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
// 	mock.ExpectQuery(regexp.QuoteMeta(
// 		`INSERT INTO "projects" ("name","user_id","created_at","updated_at") VALUES ($1,$2,$3,$4)`)).
// 		WillReturnRows(rows)
// 	mock.ExpectCommit()

// 	mock.ExpectExec("INSERT INTO \"project_user_mappings\"").
// 		WithArgs(sqlmock.AnyArg(), test2ProjectDto.UserID).
// 		WillReturnError(errors.New("insertion error"))
// 	mock.ExpectRollback()

// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		pd dto.ProjectDto
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"case1", fields{db.NewDB()}, args{*test1ProjectDto}, false},
// 		{"case2", fields{gormDB}, args{*test2ProjectDto}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			pr := &projectRepository{
// 				db: tt.fields.db,
// 			}
// 			if err := pr.CreateProject(tt.args.pd); (err != nil) != tt.wantErr {
// 				t.Errorf("projectRepository.CreateProject() error = %v, wantErr %v", err, tt.wantErr)
// 			}

// 			if tt.name == "case2" {
// 				if err := mock.ExpectationsWereMet(); err != nil {
// 					t.Errorf("there were unfulfilled expectations: %s", err)
// 				}
// 			}
// 		})
// 	}
// }
