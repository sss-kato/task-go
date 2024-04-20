package service

import (
	"errors"
	"reflect"
	"task-go/domain"
	"task-go/repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_userService_Signup(t *testing.T) {
	testUser1, _ := domain.NewUser("test11", "test11", "test11@gmail")
	testUser2, _ := domain.NewUser("test11", "test11", "test11@gmail")

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := repository.NewMockUserRepositoryIF(mockCtl)
	mock.EXPECT().RegistUser(testUser1).Do(func(user *domain.User) {
		user.Name = "test1"
		user.ID = 1
	}).Return(nil)
	mock.EXPECT().RegistUser(testUser2).Return(errors.New("test"))

	type fields struct {
		ur repository.UserRepositoryIF
	}
	type args struct {
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.UserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock}, args{testUser1}, domain.UserResponse{ID: 1, Name: "test1"}, false},
		{"case2", fields{mock}, args{testUser2}, domain.UserResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				ur: tt.fields.ur,
			}
			got, err := us.Signup(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Signup() = %v, want %v", got, tt.want)
			}
		})
	}
}
