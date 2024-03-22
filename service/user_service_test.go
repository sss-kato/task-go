package service

import (
	"errors"
	"reflect"
	"task-go/model"
	"task-go/repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_userService_Signup(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := repository.NewMockUserRepositoryIF(mockCtl)
	mock.EXPECT().CreateUser(&model.User{Name: "test1", Password: hashed("test1")}).Do(func(user *model.User) {
		user.Name = "test1"
		user.Password = hashed("test1")
		user.ID = 1
	}).Return(nil)
	mock.EXPECT().CreateUser(&model.User{Name: "test2", Password: hashed("test2")}).Return(errors.New("test"))

	type fields struct {
		ur repository.UserRepositoryIF
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.UserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock}, args{&model.User{Name: "test1", Password: "test1"}}, model.UserResponse{ID: 1, Name: "test1"}, false},
		{"case2", fields{mock}, args{&model.User{Name: "test2", Password: "test2"}}, model.UserResponse{}, true},
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
