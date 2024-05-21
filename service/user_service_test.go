package service

import (
	"errors"
	"reflect"
	"task-go/domain"
	"task-go/dto"
	"task-go/repository"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
)

func Test_userService_Signup(t *testing.T) {
	mockUser1 := domain.NewUser("test11", "test11", "test11@gmail")
	mockUser1.HashedPassword()
	ud1 := &dto.UserDto{Name: mockUser1.GetName(), Password: mockUser1.GetPassWord(), Mailadress: mockUser1.GetMailAdress()}
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := repository.NewMockUserRepositoryIF(mockCtl)
	mock.EXPECT().RegistUser(ud1).Do(func(user *dto.UserDto) {
		user.Name = "test1"
		user.ID = 1
	}).Return(nil)

	mockUser2 := domain.NewUser("test11", "test11", "test11@gmail")
	mockUser2.HashedPassword()
	ud2 := &dto.UserDto{Name: mockUser2.GetName(), Password: mockUser2.GetPassWord(), Mailadress: mockUser2.GetMailAdress()}
	mock.EXPECT().RegistUser(ud2).Return(errors.New("test"))

	type fields struct {
		ur repository.UserRepositoryIF
	}
	type args struct {
		user domain.UserIF
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.UserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock}, args{domain.NewUser("test11", "test11", "test11@gmail")}, domain.UserResponse{ID: 1, Name: "test1"}, false},
		{"case2", fields{mock}, args{domain.NewUser("test11", "test11", "test11@gmail")}, domain.UserResponse{}, true},
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

func Test_userService_Login(t *testing.T) {

	// test case1 mock
	mockUser1 := domain.NewUser("test11", "test11", "test11@gmail")
	mockUser1.HashedPassword()
	ud1 := &dto.UserDto{Name: mockUser1.GetName(), Password: mockUser1.GetPassWord()}
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mock := repository.NewMockUserRepositoryIF(mockCtl)
	mock.EXPECT().GetUser(ud1).Do(func(user *dto.UserDto) {
	}).Return(1, 1, nil)

	// test case1 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	const key = "4fe269f707e7ffdf0c772994046a4242449de81d1acef7bc2dc6588099fabec2"
	tokenString, _ := token.SignedString([]byte(key))

	type fields struct {
		ur repository.UserRepositoryIF
	}
	type args struct {
		user domain.UserIF
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", fields{mock}, args{domain.NewUser("test11", "test11", "test11@gmail")}, tokenString, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &userService{
				ur: tt.fields.ur,
			}
			got, err := us.Login(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
