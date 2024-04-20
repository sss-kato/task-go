package repository

import (
	"task-go/db"
	"task-go/domain"
	"testing"

	"gorm.io/gorm"
)

func Test_userRepository_RegistUser(t *testing.T) {

	// mockCtl := gomock.NewController(t)
	// defer mockCtl.Finish()
	// mock := domain.NewMockUserIF(mockCtl)
	// mock.EXPECT()

	testUser1, _ := domain.NewUser("test11", "test11", "test11@gmail")
	testUser2, _ := domain.NewUser("test11", "test11", "test11@gmail")

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		user *domain.UserIF
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool // エラーを返すことを確認する場合はtrue
	}{
		// TODO: Add test cases.
		{"case1", fields{db.NewDB()}, args{&testUser1}, false},
		{"case2", fields{db.NewDB()}, args{&testUser2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				db: tt.fields.db,
			}
			if err := ur.RegistUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.RegistUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func Test_userRepository_GetUser(t *testing.T) {
// 	type fields struct {
// 		db *gorm.DB
// 	}
// 	type args struct {
// 		user *model.User
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"case1", fields{db.NewDB()}, args{&model.User{Name: "test4", Password: "test4"}}, false},
// 		{"case2", fields{db.NewDB()}, args{&model.User{Name: "test6", Password: "test6"}}, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ur := &userRepository{
// 				db: tt.fields.db,
// 			}
// 			if err := ur.GetUser(tt.args.user); (err != nil) != tt.wantErr {
// 				t.Errorf("userRepository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
